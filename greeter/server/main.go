package main

import (
	"fmt"
	"github.com/alidevjimmy/grpc-test/greeter/greeterpb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (*server) Greet(cx context.Context , req *greeterpb.GreetRequest) (*greeterpb.GreetResponse , error){
	fmt.Printf("request send : %v" , req)
	firstName := req.GetGreeter().GetFirstName()
	result := "hello "+firstName

	res := &greeterpb.GreetResponse{
		Result: result,
	}

	return res , nil
}

func main() {
	lis , err := net.Listen("tcp" , ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	greeterpb.RegisterGreeterServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
