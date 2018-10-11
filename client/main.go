package main

import (
	"context"
	"log"
	"time"

	pd "../proto/product"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pd.NewProductClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetProducts(ctx, &pd.GetProductsRequest{CusAddress: "0x"})
	if err != nil {
		log.Fatalf("could not get products: %v", err)
	}
	log.Printf("Response: %d", r.Result)
}
