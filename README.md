# TODO application using gRPC 


Install protoc cmd tool 
```Install the protocol compiler plugins for Go using the following commands:

$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
Update your PATH so that the protoc compiler can find the plugins:

$ export PATH="$PATH:$(go env GOPATH)/bin"
```


Generate proto 
```
protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     todos/todo.proto 
```


## Execution 

assuming you are in project dir 

- `go mod tidy` 
- `go run server.go`
- `go run client.go` 
- 