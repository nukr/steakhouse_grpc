package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/nukr/steakhouse_grpc/pb"
	"google.golang.org/grpc"
)

var (
	endpoint = flag.String(
		"echo_endpoint", "localhost:2323", "endpoint of EchoService")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterSteakhouseHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	if err != nil {
		return err
	}
	err = http.ListenAndServe(":9091", mux)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()
	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
