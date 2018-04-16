package main

import (
	"encoding/json"
	"log"
	"time"

	pb "GoWebApi/rpcdatacontract"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:60000"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAIClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Process(ctx, &pb.Request{Question: "Ask a question"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	// log.Printf("Greeting: %s", r.GetTables())

	data, _ := json.MarshalIndent(r.GetTables(), "", "    ")
	log.Printf("%s\n", data)
}
