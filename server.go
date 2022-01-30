//server code for todo app 

package main 

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "github.com/ssp4all/grpc-todo/todos"
	"log"
	"net"

)

type Server interface {
	CreateTodo(context.Context, *pb.CreateTodoRequest) (*pb.Todo, error)
	GetAllTodos(context.Context, *pb.GetAllTodosRequest) (*pb.GetAllTodosResponse, error)
	Run()
}


type ToDoServer struct {
	pb.UnimplementedTodoServiceServer  //for forward compatibility
	// todo_list *pb.GetAllTodoResponse
}

func NewToDoServer() *ToDoServer {
	return &ToDoServer{
		// todo_list: &pb.Todo{},
	}
}

func (s *ToDoServer) Run() error {
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

	return nil
}

func (s *ToDoServer) GetAllTodos(ctx context.Context, req *pb.GetAllTodosRequest) (*pb.GetAllTodosResponse, error) {
	log.Println("GetAllTodos")
	toDoList := pb.GetAllTodosResponse{
		Todos: []*pb.Todo{
			&pb.Todo{
				Id: 1,
				Title: "Laundry",
				Text: "Do laundry",
			},
			&pb.Todo{
				Id: 2,
				Title: "Study",
				Text: "Do study",
			},
		},
	}
	log.Println("GetAllTodos: ", toDoList)
	return &toDoList, nil
}


// CreateTodo implements TodoService.CreateTodo
func (s *ToDoServer) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.Todo, error) {
	//log request content
	log.Printf("\nCreateTodo: %s\n%s", req.Title, req.Text)

	// id := len(s.todo_list.ToDos) + 1 //assign id to new todo
	id := int32(11)
	//append todo into allToDos
	// s.todo_list = append(s.todo_list, &pb.ToDo{
	// 	Id: id,
	// 	Title: req.Title,
	// 	Text: req.Text,
	// })


	return &pb.Todo{
		Id: id,
		Title: req.Title,
		Text: req.Text,
	}, nil
}


func main(){
	//init server
	server := NewToDoServer()
	//start server
	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	
}