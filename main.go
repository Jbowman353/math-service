package main

import (
	"fmt"
	// "net/http"
	"net"
	"log"
	"google.golang.org/grpc"
	"context"
	pb "github.com/Jbowman353/math-service/math"
)

const (
	port = ":55555"
)

type server struct {
	pb.UnimplementedMathServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Add(ctx context.Context, in *pb.Operands) (*pb.MathRes, error) {
	toAdd := in.GetOperands();
	var sum int64 = 0
	for _, val := range toAdd {
		sum += val
	}
	return &pb.MathRes{ Result: sum }, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMathServiceServer(s, &server{})
	fmt.Println("Serving Now...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return
}