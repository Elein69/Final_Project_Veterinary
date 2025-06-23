package main

import (
	"context"
	"log"
	"net"

	"audit_service/proto"

	"google.golang.org/grpc"
)

type auditServer struct {
	proto.UnimplementedAuditServiceServer
}

func (s *auditServer) LogEvent(ctx context.Context, req *proto.AuditRequest) (*proto.AuditResponse, error) {
	log.Printf("AUDIT LOG -> UserID: %s, Action: %s, Timestamp: %s\n", req.UserId, req.Action, req.Timestamp)
	return &proto.AuditResponse{Status: "Success"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8087")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterAuditServiceServer(grpcServer, &auditServer{})

	log.Println("✅ gRPC server started on port 8087...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
