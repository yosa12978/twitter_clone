package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/yosa12978/twitter/user-api/db"
	"github.com/yosa12978/twitter/user-api/handlers"
	"github.com/yosa12978/twitter/user-api/logging"
)

func main() {
	logger := logging.New("main")
	if err := godotenv.Load(); err != nil {
		logger.Fatalf(err.Error())
	}
	ctx := context.Background()

	db.GetDB(ctx)

	addr := os.Getenv("ADDR")
	if addr == "" {
		logger.Fatalf("listening address must be specified in env args")
	}
	server := http.Server{
		Addr:           addr,
		MaxHeaderBytes: 1 << 10,
		Handler:        handlers.NewRouter(),
	}
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Fatalf(err.Error())
	}
	go server.Serve(listener)
	logger.Printf("Listening @ %s", addr)

	out := make(chan os.Signal, 1)
	signal.Notify(out, os.Interrupt, syscall.SIGTERM)
	sig := <-out
	logger.Printf("Program halted: %s", sig.String())
}
