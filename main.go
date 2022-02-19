package main

import (
	"encoding/json"
	"github.com/luizaugustoventura/aula-grpc-go/grpc/pb"
	"github.com/luizaugustoventura/aula-grpc-go/grpc/service"
	"github.com/luizaugustoventura/aula-grpc-go/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

var ProductList = model.NewProducts() 

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Print("Erro ao se conectar")
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	productService := service.NewProductGrpcService()
	productService.Products = ProductList
	pb.RegisterProductServiceServer(grpcServer, productService)

	go grpcServer.Serve(lis)

	http.HandleFunc("/products", HandleProducts)
	http.ListenAndServe(":8080", nil)
}

func HandleProducts(w http.ResponseWriter, r *http.Request) {
	productsJson, err := json.Marshal(ProductList)
	if err != nil {
		return
	}
	w.Write([]byte(productsJson))
}