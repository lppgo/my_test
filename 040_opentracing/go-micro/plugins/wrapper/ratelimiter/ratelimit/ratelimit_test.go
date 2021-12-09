package ratelimit

import (
	"fmt"
	"testing"
	"time"

	"context"

	"github.com/juju/ratelimit"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/errors"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
	"go-micro.dev/v4/server"
	"go-micro.dev/v4/transport"
)

type testHandler struct{}
type TestRequest struct{}
type TestResponse struct{}

func (t *testHandler) Method(ctx context.Context, req *TestRequest, rsp *TestResponse) error {
	return nil
}

func TestRateClientLimit(t *testing.T) {
	// setup
	r := registry.NewMemoryRegistry()
	s := selector.NewSelector(selector.Registry(r))
	tr := transport.NewMemoryTransport()
	testRates := []int{1, 10, 20}

	for _, limit := range testRates {
		b := ratelimit.NewBucketWithRate(float64(limit), int64(limit))

		c := client.NewClient(
			// set the selector
			client.Selector(s),
			client.Transport(tr),
			// add the breaker wrapper
			client.Wrap(NewClientWrapper(b, false)),
		)

		req := c.NewRequest(
			"test.service",
			"Test.Method",
			&TestRequest{},
			client.WithContentType("application/json"),
		)
		rsp := TestResponse{}

		for j := 0; j < limit; j++ {
			err := c.Call(context.TODO(), req, &rsp)
			e := errors.Parse(err.Error())
			if e.Code == 429 {
				t.Errorf("Unexpected rate limit error: %v", err)
			}
		}

		err := c.Call(context.TODO(), req, rsp)
		e := errors.Parse(err.Error())
		if e.Code != 429 {
			t.Errorf("Expected rate limit error, got: %v", err)
		}
	}
}

func TestRateServerLimit(t *testing.T) {
	// setup
	testRates := []int{1, 5, 6, 10}

	for _, limit := range testRates {
		r := registry.NewMemoryRegistry()
		b := broker.NewMemoryBroker()
		tr := transport.NewMemoryTransport()
		_ = b
		s := selector.NewSelector(selector.Registry(r))

		br := ratelimit.NewBucketWithRate(float64(limit), int64(limit))
		c := client.NewClient(client.Selector(s), client.Transport(tr))

		name := fmt.Sprintf("test.service.%d", limit)

		srv := server.NewServer(
			server.Name(name),
			// add registry
			server.Registry(r),
			server.Transport(tr),
			// add broker
			//server.Broker(b),
			// add the breaker wrapper
			server.WrapHandler(NewHandlerWrapper(br, false)),
		)

		type Test struct {
			*testHandler
		}

		srv.Handle(
			srv.NewHandler(&Test{new(testHandler)}),
		)

		if err := srv.Start(); err != nil {
			t.Fatalf("Unexpected error starting server: %v", err)
		}
		req := c.NewRequest(name, "Test.Method", &TestRequest{}, client.WithContentType("application/json"))
		rsp := TestResponse{}

		for j := 0; j < limit; j++ {
			if err := c.Call(context.TODO(), req, &rsp); err != nil {
				t.Fatalf("Unexpected request error: %v", err)
			}
		}

		err := c.Call(context.TODO(), req, &rsp)
		if err == nil {
			t.Fatalf("Expected rate limit error, got nil: rate %d, err %v", limit, err)
		}

		e := errors.Parse(err.Error())
		if e.Code != 429 {
			t.Fatalf("Expected rate limit error, got %v", err)
		}

		srv.Stop()

		// artificial test delay
		time.Sleep(500 * time.Millisecond)
	}
}
