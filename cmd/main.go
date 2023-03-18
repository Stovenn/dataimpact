package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/stovenn/dataimpact/pkg/http"
	"github.com/stovenn/dataimpact/pkg/mongo"
)

func main() {
	infoLogger := log.New(os.Stdin, "[INFO]", log.LstdFlags)
	errLogger := log.New(os.Stderr, "[ERROR]", log.LstdFlags)

	client := mongo.InitClient()
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	store := mongo.NewUserStore(client)

	server := http.NewServer(store, infoLogger, errLogger)

	go func() {
		fmt.Printf("Server listening on port 8080\n")
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errLogger.Fatalf("an error occured on the server: %v", err)
		}
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	infoLogger.Println("received terminate, graceful shutdown", sig)
	if err := server.Shutdown(ctx); err != nil {
		errLogger.Printf("error on server shutdown: %v", err)
	}

	server.ListenAndServe()
}