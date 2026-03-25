package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/lorekeeper/grpc/pb"
	"github.com/lorekeeper/grpc/server"
)

const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterLoreKeeperRAGServer(grpcServer, &server.RagServer{})

	reflection.Register(grpcServer)

	log.Printf("LoreKeeper RAG Engine started on port %s", port)
	log.Printf("   Awaiting queries from the Backend...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
