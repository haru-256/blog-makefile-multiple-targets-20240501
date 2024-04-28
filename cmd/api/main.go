package main

import (
	"log/slog"
	"net/http"
	"os"
)

var logger *slog.Logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func hello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello, World!"))
	if err != nil {
		logger.Error("Write Error", "Error", err)
	}
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Get Request", "RequestURI", r.RequestURI)
		h.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", log(hello))
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		logger.Error("Server Error", "Error", err)
	}
}
