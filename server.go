package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/ssp4all/grpc-todo/todos"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func NewToDoServer() *ToDoServer {
	return &ToDoServer{
		todo_list: &pb.GetAllTodosResponse{},
	}
}

type ToDoServer struct {
	pb.UnimplementedTodoServiceServer
	todo_list *pb.GetAllTodosResponse
}

func (server *ToDoServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	return s.Serve(lis)
}

func (server *ToDoServer) CreateTodo(ctx context.Context, in *pb.CreateTodoRequest) (*pb.Todo, error) {
	log.Printf("Received: %v", in.GetTitle())
	var todo_id = int32(rand.Intn(100))
	created_todo := &pb.Todo{Title: in.GetTitle(), Text: in.GetText(), Id: todo_id}
	server.todo_list.Todos = append(server.todo_list.Todos, created_todo)
	return created_todo, nil
}

func (server *ToDoServer) GetAllTodos(ctx context.Context, in *pb.GetAllTodosRequest) (*pb.GetAllTodosResponse, error) {
	return server.todo_list, nil
}

func main() {
	var todo_server *ToDoServer = NewToDoServer()
	if err := todo_server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
