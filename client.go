//grpc client code

package main

import (
	"context"
	// "time"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
	pb "github.com/ssp4all/grpc-todo/todos"
	"log"
	// "net"
)

func main() {
	//dial to port 50051
	log.Println("Dialng to port 50051")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTodoServiceClient(conn)
	//call CreateTodo
	log.Println("CreateTodo")
	r, err := c.CreateTodo(context.Background(), &pb.CreateTodoRequest{
		Title: "Laundry",
		Text: "Do laundry",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Created todo: %s\n%s", r.Title, r.Text)

	r, err := c.CreateTodo(context.Background(), &pb.CreateTodoRequest{
		Title: "Study",
		Text: "Do study",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	//log response
	log.Printf("Created todo: %s\n%s", r.Title, r.Text)


	//use getToDo by id 1 
	log.Println("GetToDo")
	r, err = c.GetToDo(context.Background(), &pb.GetToDoRequest{
		Id: 1,
	})
	if err != nil {
		log.Fatalf("could not found todo by id %v", err)
	}
	log.Printf("Found todo: %s\n%s", r.Title, r.Text)
	

}
