package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os/signal"
	"syscall"
	"user_service/infra/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.Load(ctx)
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		log.Fatal(err)
	}

	svr := setupServer(ctx)
	go func() {
		<-ctx.Done()
		svr.GracefulStop()
	}()

	if err := svr.Serve(l); err != nil {
		log.Fatal(err)
	}
}

func setupServer(ctx context.Context) *grpc.Server {
	// todo := infra.NewTodoRepository()
	svr := grpc.NewServer()
	// pb.RegisterTodoApiServer(
	// 	svr,
	// 	ui.NewGRPCService(
	// 		app.NewTodoService(todo),
	// 	),
	// )
	reflection.Register(svr)

	return svr
}
