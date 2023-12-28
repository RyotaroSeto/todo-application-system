package handlers

import (
	"context"
	"fmt"
	"gateway/configs"
	"net/http"
)

func NewHandler(ctx context.Context, cfg *configs.Config) (http.Handler, error) {
	mux := http.NewServeMux()

	mux.HandleFunc("/todo", corsMiddleware(todoHandler))
	return mux, nil
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprint(w, "GET hello!\n")
	case "POST":
		fmt.Fprint(w, "POST hello!\n")
	case "DELETE":
		fmt.Fprint(w, "DELETE hello!\n")
	default:
		fmt.Fprint(w, "Method not allowed.\n")
	}
}
