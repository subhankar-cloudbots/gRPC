package main

import (
	"context"
	"log"

	pb "github.com/subhankar-cloudbots/grpc-basic/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Tanmoy",
	})

	if err != nil {
		log.Fatalln("Could not greet")
	}

	log.Printf("Greeting: %s\n", res.Result)
}
