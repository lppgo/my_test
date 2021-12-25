package main

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"time"

	pb "github.com/asim/go-micro/examples/v4/stream/grpc/proto"
	"github.com/asim/go-micro/plugins/client/grpc/v4"
	"github.com/uber/jaeger-client-go"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	wrapperTrace "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	"github.com/opentracing/opentracing-go"

	traceconfig "github.com/asim/go-micro/examples/v4/stream/grpc/tracer"
)

var (
	microServerName = "RouteGuide.Server"
	microClientName = "RouteGuide.Client"

	tracerServiceName = "RouteGuide-Client"
	// tracerSamplerType  = "const"
	// tracerSamplerParam = float64(1)
	// tracerAgentAddr = "localhost:6831"
)

func main() {
	// tracer
	var tracer opentracing.Tracer
	tracerCfg := traceconfig.TracerCfg{
		ServiceName: tracerServiceName,
		// AgentAddr:    "localhost:6831",
		AgentAddr:    "172.17.0.1:6831",
		Disable:      false,
		SamplerType:  jaeger.SamplerTypeConst,
		SamplerParam: 1,
	}
	tracer, close := traceconfig.TraceInit(tracerCfg)
	defer close.Close()
	opentracing.SetGlobalTracer(tracer)

	// micro service
	srv := micro.NewService(
		micro.Client(grpc.NewClient()),
		micro.Name(microClientName),
		micro.WrapClient(wrapperTrace.NewClientWrapper(opentracing.GlobalTracer())),
	)
	srv.Init()
	client := pb.NewRouteGuideService(microServerName, srv.Client())

	// request
	// Looking for a valid feature
	fmt.Println("---------get a valid feature----------")
	printFeature(client, &pb.Point{Latitude: 409146138, Longitude: -746188906}) //nolint:gomnd

	// Feature missing. 测试tag error=true
	fmt.Println("---------get a invalid feature----------")
	printFeature(client, &pb.Point{Latitude: 0, Longitude: 0})

	// Looking for features between 40, -75 and 42, -73.
	fmt.Println("---------get features by Rectangle----------")
	printFeatures(client, &pb.Rectangle{
		Lo: &pb.Point{Latitude: 400000000, Longitude: -750000000}, //nolint:gomnd
		Hi: &pb.Point{Latitude: 420000000, Longitude: -730000000}, //nolint:gomnd
	})
	// RecordRoute
	fmt.Println("---------invoke RecordRoute ... ----------")
	runRecordRoute(client)

	// RouteChat
	fmt.Println("---------invoke RouteChat ... ----------")
	runRouteChat(client)
}

// printFeature gets the feature for the given point.
func printFeature(client pb.RouteGuideService, point *pb.Point) {
	logger.Infof("Getting feature for point (%d, %d)", point.Latitude, point.Longitude)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	feature, err := client.GetFeature(ctx, point)
	if err != nil {
		logger.Errorf("getfeature error : %v", err)
		return
	}
	logger.Info(feature)
}

// printFeatures lists all the features within the given bounding Rectangle.
func printFeatures(client pb.RouteGuideService, rect *pb.Rectangle) {
	logger.Infof("Looking for features within %v", rect)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.ListFeatures(ctx, rect)
	if err != nil {
		logger.Errorf("get list feature error : %v", err)
		return
	}
	// IMPORTANT: do not forgot to close stream
	defer stream.Close()
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Error(err)
			break
		}
		logger.Infof("Feature: name: %q, point:(%v, %v)", feature.GetName(),
			feature.GetLocation().GetLatitude(), feature.GetLocation().GetLongitude())
	}
}

// runRecordRoute sends a sequence of points to server and expects to get a RouteSummary from server.
func runRecordRoute(client pb.RouteGuideService) {
	// Create a random number of random points
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pointCount := int(r.Int31n(100)) + 2 // Traverse at least two points
	var points []*pb.Point
	for i := 0; i < pointCount; i++ {
		points = append(points, randomPoint(r))
	}
	logger.Infof("Traversing %d points.", len(points))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.RecordRoute(ctx)
	if err != nil {
		logger.Errorf("record route error : %v", err)
		return
	}
	// IMPORTANT: do not forgot to close stream
	defer stream.Close()
	for _, point := range points {
		if err := stream.Send(point); err != nil {
			logger.Errorf("stream send error : %v", err)
		}
	}
	if err := stream.CloseSend(); err != nil {
		logger.Errorf("stream closeSend error : %v", err)
		return
	}
	summary := pb.RouteSummary{}
	if err := stream.RecvMsg(&summary); err != nil {
		logger.Errorf("stream recvMsg error : %v", err)
		return
	}
	logger.Infof("Route summary: %v", &summary)
}

// runRouteChat receives a sequence of route notes, while sending notes for various locations.
func runRouteChat(client pb.RouteGuideService) {
	notes := []*pb.RouteNote{
		{Location: &pb.Point{Latitude: 0, Longitude: 1}, Message: "First message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 2}, Message: "Second message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 3}, Message: "Third message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 1}, Message: "Fourth message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 2}, Message: "Fifth message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 3}, Message: "Sixth message"},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.RouteChat(ctx)
	if err != nil {
		logger.Errorf("routeChat error : %v", err)
		return
	}
	// IMPORTANT: do not forgot to close stream
	defer stream.Close()
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				break
			}
			if err != nil {
				logger.Fatal(err)
			}
			logger.Infof("Got message %s at point(%d, %d)", in.Message, in.Location.Latitude, in.Location.Longitude)
		}
	}()
	for _, note := range notes {
		if err := stream.Send(note); err != nil {
			logger.Error(err)
		}
	}
	if err := stream.CloseSend(); err != nil {
		logger.Errorf("stream closeSend error : %v", err)
		return
	}
	<-waitc
}

func randomPoint(r *rand.Rand) *pb.Point {
	lat := (r.Int31n(180) - 90) * 1e7
	long := (r.Int31n(360) - 180) * 1e7
	return &pb.Point{Latitude: lat, Longitude: long}
}
