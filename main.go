package main

import (
	"log"
	"net"

	db "go-todolist/DB"
	"go-todolist/service"
	pb "go-todolist/service/todo"

	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTodoListServer(s, &service.Server{
		GrpcService: &db.SVC,
	})
	log.Printf("Starting gRPC listener on port " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
