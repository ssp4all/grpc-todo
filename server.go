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

type Server interface {
	CreateTodo(context.Context, *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error)
	GetToDo(context.Context, *pb.GetToDoRequest) (*pb.GetToDoResponse, error)
}



type ToDoServer struct {
	pb.UnimplementedTodoServiceServer
	todo_list *pb.allToDos
}


// CreateTodo implements TodoService.CreateTodo
func (s *ToDoServer) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	fmt.Println("CreateTodo")
	//log request content
	log.Println("CreateTodo: %s\n%s", req.Title, req.Text)

	id := len(s.todo_list.ToDos) + 1 //assign id to new todo

	//append todo into allToDos
	s.todo_list = append(s.todo_list, &pb.ToDo{
		Id: id,
		Title: req.Title,
		Text: req.Text,
	})


	return &pb.CreateTodoResponse{
		Id: id,
		Title: req.Title,
		Text: req.Text,
	}, nil
}

//implement getToDo
func (s *ToDoServer) GetToDo(ctx context.Context, req *pb.GetToDoRequest) (*pb.GetToDoResponse, error) {
	//log inside getToDo
	log.Println("GetToDo func")

	//find todo by id in todo_list
	for _, todo := range s.todo_list.ToDos {
		if todo.Id == req.Id {
			return &pb.GetToDoResponse{
				Id: todo.Id,
				Title: todo.Title,
				Text: todo.Text,
			}, nil
		}
	}

	//if not found, return error
	return nil, fmt.Errorf("todo not found")

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
	pb.RegisterTodoServiceServer(server, &ToDoServer{})
	log.Println("server started")

	//start server
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	
}