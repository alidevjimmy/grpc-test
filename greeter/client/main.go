package main

import (
	"context"
	"github.com/alidevjimmy/grpc-test/greeter/greeterpb"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn , err := grpc.Dial("localhost:50051" , grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	client := greeterpb.NewGreeterServiceClient(conn)
	doUnary(client)
}


func doUnary(client greeterpb.GreeterServiceClient) {
	fmt.Println("Starting send request...")
	greet := &greeterpb.GreetRequest{
		Greeter: &greeterpb.Greet{
			FirstName: "ali",
			LastName: "hamrani",
		},
	}
	resp ,err := client.Greet(context.Background() , greet)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp.GetResult())
}