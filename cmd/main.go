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
	"github.com/stovenn/dataimpact/pkg/util"
)

func main() {
	util.CreateDataDirIfNotExists()
	config, err := util.SetupConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v\n", err)
	}

	infoLogger := log.New(os.Stdin, "[INFO]", log.LstdFlags)
	errLogger := log.New(os.Stderr, "[ERROR]", log.LstdFlags)

	mongo.InitMongoStore(config.DBUri, config.DBName)
	defer func() {
		if err := mongo.S.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	server := http.NewServer(mongo.S, infoLogger, errLogger, config)

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
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	infoLogger.Println("received terminate, graceful shutdown", sig)
	if err := server.Shutdown(ctx); err != nil {
		errLogger.Printf("error on server shutdown: %v", err)
	}

	server.ListenAndServe()
}
