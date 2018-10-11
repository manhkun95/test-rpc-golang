package main

import (
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	pd "../proto/product"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const port = ":9000"

type server struct{}

func (s *server) GetProducts(ctx context.Context, in *pd.GetProductsRequest) (*pd.GetProductsResponse, error) {
	return &pd.GetProductsResponse{
		Result: 1,
		Count:  0,
		Page:   1,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	log.Printf("Server is listen on post: %s", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pd.RegisterProductServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
