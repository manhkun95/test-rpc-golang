package main

import (
	"context"
	"log"
	"time"
	"github.com/golang/protobuf/jsonpb"

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
	log.Printf("----------------CREATE PRODUCT------------------------")
	resCreate, err := c.CreateProduct(ctx, &pd.CreateProductRequest{Name: "HeadPhone2", Description: ""})
	log.Printf("Create product result %d", resCreate.Result)
	log.Printf("----------------GET LIST PRODUCT----------------------")
	resGet, err := c.GetProducts(ctx, &pd.GetProductsRequest{CusAddress: "0x"})
	if err != nil {
		log.Fatalf("could not get products: %v", err)
	}
	for i := range resGet.Products {
		m := jsonpb.Marshaler{}
		js, _ := m.MarshalToString(resGet.Products[i])
        log.Printf(js)
    }
}
