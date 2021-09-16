package service

import (
	"context"
	"fmt"

	db "go-todolist/DB"
	pb "go-todolist/service/todo"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	TodoDB *db.MongoClient
}

func (s *Server) AddTodoList(ctx context.Context, in *pb.Todo) (*pb.TodoID, error) {
	result := s.TodoDB.InsertTodo(in.NickName, in.ToDo)
	return &pb.TodoID{Value: fmt.Sprintf("%x", result.InsertedID)}, status.New(codes.OK, "").Err()
}
