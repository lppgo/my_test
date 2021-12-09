package web_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/web"
)

func TestWeb(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println("Test nr", i)
		testFunc()
	}
}

func testFunc() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*250)
	defer cancel()

	s := micro.NewService(
		micro.Name("test"),
		micro.Context(ctx),
		micro.HandleSignal(false),
		micro.Flags(
			&cli.StringFlag{
				Name: "test.timeout",
			},
			&cli.BoolFlag{
				Name: "test.v",
			},
			&cli.StringFlag{
				Name: "test.run",
			},
			&cli.StringFlag{
				Name: "test.testlogfile",
			},
		),
	)
	w := web.NewService(
		web.MicroService(s),
		web.Context(ctx),
		web.HandleSignal(false),
	)
	//s.Init()
	//w.Init()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := s.Run()
		if err != nil {
			logger.Errorf("micro run error: %v", err)
		}
	}()
	go func() {
		defer wg.Done()
		err := w.Run()
		if err != nil {
			logger.Errorf("web run error: %v", err)
		}
	}()

	wg.Wait()
}
