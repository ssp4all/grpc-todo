//server code for todo app 

package main 

import (
	"context"
	"fmt"
	// "time"
	//import grpc
	"google.golang.org/grpc"
	//import protobuf
	pb "github.com/ssp4all/grpc-todo/todos"
	"log"
	"net"

)

type Server struct{
	pb.UnimplementedTodoServiceServer //to avoid *Server does not implement todo.TodoServiceServer (missing todo.mustEmbedUnimplementedTodoServiceServer method)
	
}

// CreateTodo implements TodoService.CreateTodo
func (s *Server) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	fmt.Println("CreateTodo")
	return &pb.CreateTodoResponse{
		Title: req.Title,
		Text: req.Text,
	}, nil
}

func main(){
	//listen to port 50051
	fmt.Println("listening on port 50051")
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//init server
	server := grpc.NewServer()
	//register service
	pb.RegisterTodoServiceServer(server, &Server{})
	log.Println("server started")

	//start server
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	
}