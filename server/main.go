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

func (s *server) CreateProduct(ctx context.Context, in *pd.CreateProductRequest) (*pd.CreateProductResponse, error) {
	log.Printf("Create product with name: %s", in.Name)
	return &pd.CreateProductResponse{
		Result: 1,
	}, nil
}

func (s *server) GetProducts(ctx context.Context, in *pd.GetProductsRequest) (*pd.GetProductsResponse, error) {
	return &pd.GetProductsResponse{
		Result: 1,
		Products: []*pd.Products{
			{
				Id: 1,
				Name: "Head Phone",
				Description: "This is a device with...",
				Detail: "...",
				Imgs: "123",
				Properties: "properties",
				Price: 20000,
				TotalNumber: 6,
			},
			{
				Id: 2,
				Name: "Head Phone 2",
				Description: "This is a device with...",
				Detail: "...",
				Imgs: "123",
				Properties: "properties",
				Price: 20000,
				TotalNumber: 6,
			},
	    },
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
