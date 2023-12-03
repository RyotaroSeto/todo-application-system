package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os/signal"
	"syscall"

	"auth_service/infra/config"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	repo, err := config.Load(ctx)
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", repo.Port))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(l)

	// svr := setupServer(repo)
	// go func() {
	// 	<-ctx.Done()
	// 	svr.GracefulStop()
	// }()

	// if err := svr.Serve(l); err != nil {
	// 	log.Fatal(err)
	// }
}

// func setupServer(repo *config.Config) interface{} {
// 	return nil
// }
