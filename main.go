package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/AlexeyBorisovets/USER/internal/model"
	"github.com/AlexeyBorisovets/USER/internal/repository"
	"github.com/AlexeyBorisovets/USER/internal/server"
	"github.com/AlexeyBorisovets/USER/internal/service"
	pb "github.com/AlexeyBorisovets/USER/proto"
	"github.com/caarlos0/env/v6"
	"google.golang.org/grpc"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		defer log.Fatalf("error while listening port: %e", err)
	}
	fmt.Println("Server successfully started on port :8080...")

	key := []byte("super-key")
	cfg := model.Config{JwtKey: key}
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("failed to start service, %e", err)
	}

	poolP, err := pgxpool.Connect(context.Background(), "postgres://postgres:123@localhost:5432/Test")
	if err != nil {
		log.Fatalf("bad connection with postgresql: %v", err)
	} else {
		fmt.Println("DB successfully connect...")
	}
	rP := repository.PRepository{Pool: poolP}
	defer func() {
		poolP.Close()
	}()

	ns := grpc.NewServer()
	newService := service.NewService(&rP, cfg.JwtKey)
	srv := server.NewServer(newService)
	pb.RegisterUSERServer(ns, srv)

	if err = ns.Serve(listen); err != nil {
		defer log.Fatalf("error while listening server: %e", err)
	}

	// testusers := []model.User{
	// 	{
	// 		ID:           "id1",
	// 		Name:         "Alexey",
	// 		Password:     "123",
	// 		Balance:      50,
	// 		RefreshToken: "refresh1",
	// 		UserType:     "consumer",
	// 	},
	// 	{
	// 		ID:           "id2",
	// 		Name:         "Vlad",
	// 		Password:     "123",
	// 		Balance:      55,
	// 		RefreshToken: "refresh2",
	// 		UserType:     "vendor",
	// 	},
	// }
}
