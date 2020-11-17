package main

import (
	"context"
	"log"
	"time"
	"fmt"

	"google.golang.org/grpc"
	pb "github.com/Jbowman353/math-service/math"
)

const (
	address = "localhost:55555"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMathServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	a := []int64{1,2,3}

	addRes, err := c.Add(ctx, &pb.Operands{Operands: a})
	
	fmt.Println(addRes.GetResult())

}