package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"proto-workshop/client"
	"proto-workshop/controller"
	"proto-workshop/generated/proto"
	"proto-workshop/service"
	"syscall"

	loglib "bitbucket.bri.co.id/nds/nds-go-module-logger/lib"
	"google.golang.org/grpc"
)

func main() {
	logger := loglib.NewLib()
	logger.Init("NDS", "V1.0.0")
	cl := client.NewClientStruct(logger)
	ctr := controller.NewController(logger, cl)
	svc := service.NewService(logger, ctr)

	grpcServer := grpc.NewServer()

	proto.RegisterWorkshopServiceServer(grpcServer, svc)
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "", 8081))
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start Ln Service gRPC server: %v", err)
		}
	}()

	fmt.Printf("gRPC server is running at %s:%d\n", "", 8081)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal := <-c
	log.Fatalf("process killed with signal: %v\n", signal.String())
}
