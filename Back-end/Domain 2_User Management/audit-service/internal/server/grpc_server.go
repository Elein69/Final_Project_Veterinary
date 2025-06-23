package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartGRPCServer() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	// proto.RegisterAuditServiceServer(grpcServer, &yourService{}) // si ya implementaste esto

	log.Println("gRPC server is running on port 50051")
	return grpcServer.Serve(lis)
}
