package main

import (
	"context"
	"fmt"
	"gateway/configs"
	"gateway/handlers"
	"gateway/server"
	"log"
	"net"
	"os"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminated server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := configs.LoadConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		return err
	}

	mux, err := handlers.NewHandler(ctx, cfg)
	if err != nil {
		return err
	}

	s := server.NewServer(l, mux)
	return s.Run(ctx)
}
