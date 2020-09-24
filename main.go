package main

import (
	"log"
	"net"

	"github.com/livspaceeng/zeebe-gopro/configs"
	"github.com/livspaceeng/zeebe-gopro/pkg/gateway"
	"github.com/zeebe-io/zeebe/clients/go/pkg/pb"

	"google.golang.org/grpc"
)

const (
	serverPort  = ":5050"
	gatewayAddr = "0.0.0.0:26500"
)

func main() {
	config, err := configs.GetConfig()
	var address, port string
	if err != nil {
		port = serverPort
		address = gatewayAddr
	} else {
		port = config.Server.Port
		address = config.Zeebe.Host
	}

	log.Println("Listening on port:", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Initializing connection to gateway at:", address)
	err = gateway.Init(address)
	if err != nil {
		log.Fatalf("Failed to initialize Zeebe client: %v", err)
	}

	log.Println("Starting grpc server...")
	s := grpc.NewServer()
	log.Println("grpc server started. Registering gateway server...")
	pb.RegisterGatewayServer(s, new(gateway.GatewayServerImpl))
	log.Println("Gateway server registered.")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
