package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"tradingitem.server/core"
	protopack "tradingitem.server/generated_proto"
)

func startServer() {

	osSignalChan := make(chan os.Signal, 1)
	signal.Notify(osSignalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	grpcServer := grpc.NewServer()
	protopack.RegisterAccountServicesServer(grpcServer, core.NewServer())

	grpcServerEndpoint := flag.String("grpc-server-endpoint", "localhost:8000", "grpc server endpoint")
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	ctx := context.Background()
	err := protopack.RegisterAccountServicesHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)

	httpMux := http.NewServeMux()
	httpMux.Handle("/", mux)

	httpServer := http.Server{
		Addr:    ":8001",
		Handler: httpMux,
	}

	if err != nil {
		log.Fatalf("Err occured while registing handler from end point")
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		fmt.Println("Reverse proxy is running")
		defer wg.Done()
		err := httpServer.ListenAndServe()
		if err != nil {
			log.Fatalf("Error while http server listening and serving")
		}
	}()

	go func() {
		fmt.Println("Grpc server is running")
		defer wg.Done()
		lis, err := net.Listen("tcp", "0.0.0.0:8000")
		if err != nil {
			log.Fatalf("Err occured while start a listener")
		}
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatalf("Err occured while grpc server was serving the listener")
		}
	}()

	//--------------------------------
	// Graceful Shutdown
	//--------------------------------
	<-osSignalChan

	ctx = context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	grpcServer.GracefulStop()
	err = httpServer.Shutdown(ctx)

	if err != nil {
		panic(err)
	}

	wg.Wait()

}


func main() {

	startServer()

}