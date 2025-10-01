package main

import (
	"log"
	"net"

	cartpb "github.com/e-commerce-microservice/shopping-cart/gen"
	"github.com/e-commerce-microservice/shopping-cart/internals/config"
	"github.com/e-commerce-microservice/shopping-cart/internals/db"
	"github.com/e-commerce-microservice/shopping-cart/internals/repo"
	"github.com/e-commerce-microservice/shopping-cart/internals/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config := config.NewConfig()

	db, err := db.ConnectToDB(config)

	if err != nil {
		panic(err)
	}
	cartRepo := repo.CreateRepo(db)

	cartService := service.NewCartService(cartRepo)

	lis, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	s := grpc.NewServer()

	cartpb.RegisterShoppingCartServiceServer(s, cartService)

	reflection.Register(s)

	log.Printf("Cart server listening on %s", config.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("serve: %v", err)
	}

}
