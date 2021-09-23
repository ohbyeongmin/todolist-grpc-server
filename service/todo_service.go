package service

import (
	"context"

	db "go-todolist/DB"
	pb "go-todolist/service/todo"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	GrpcService *db.Service
}

func (s *Server) AddTodoList(ctx context.Context, in *pb.Todo) (*pb.TodoID, error) {
	result := s.GrpcService.DBservice.InsertTodo(in.NickName, in.ToDo)
	return &pb.TodoID{Value: result}, status.New(codes.OK, "").Err()
}
