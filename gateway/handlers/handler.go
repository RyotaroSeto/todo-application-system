package handlers

import (
	"context"
	"gen/go/todo"
	"net/http"
	"net/http/httputil"
	"net/url"
	"pkg/errors"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcServerAddress = "todo_service:8080"
	docsServerAddress = "http://docs-server:8080"
)

func NewHandler(ctx context.Context) (http.Handler, error) {
	grpcGateway := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := todo.RegisterTodoApiHandlerFromEndpoint(ctx, grpcGateway, grpcServerAddress, opts); err != nil {
		return nil, err
	}

	docsURL, err := url.Parse(docsServerAddress)
	if err != nil {
		return nil, errors.Newf(errors.InternalServerError, "failed to parse docsServerAddress=%v:%s", docsServerAddress, err)
	}
	docsProxy := httputil.NewSingleHostReverseProxy(docsURL)

	mux := http.NewServeMux()

	mux.Handle("/docs/", docsProxy)
	mux.Handle("/", corsMiddleware(grpcGateway))
	return mux, nil
}
