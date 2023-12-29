package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os/signal"
	"syscall"
	"todo_service/app"
	"todo_service/infra"
	pb "todo_service/proto"
	"todo_service/ui"

	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	repo, err := infra.LoadConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", repo.Port))
	if err != nil {
		log.Fatal(err)
	}

	svr := setupServer()
	go func() {
		<-ctx.Done()
		svr.GracefulStop()
	}()

	if err := svr.Serve(l); err != nil {
		log.Fatal(err)
	}
}

func setupServer() *grpc.Server {
	todo := infra.NewTodoRepository()
	svr := grpc.NewServer()
	pb.RegisterTodoServiceServer(
		svr,
		ui.NewGRPCService(
			app.NewTodoService(todo),
		),
	)
	return svr
}
