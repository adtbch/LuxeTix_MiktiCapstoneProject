package main

import (
	"context"
	"fmt"
	"log"
	"luxe/config"
	"luxe/internal/builder"
	"luxe/pkg/database"
	"luxe/pkg/server"
	"os"
	"os/signal"
	"time"
)

func main() {
	// load configuration via env
	cfg, err := config.NewConfig(".env")
	checkError(err)
	fmt.Println(cfg)
	// init and start db
	db, err := database.InitDatabase(cfg.PgConfig)
	checkError(err)

	publicRoutes := builder.BuildPublicRoutes(cfg, db)

	fmt.Println("db connected")

	// init server
	srv := server.NewServer(cfg, publicRoutes)
	//start server
	runServer(srv, cfg.PORT)
	waitForShutdown(srv)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func runServer(srv *server.Server, port string) {
	go func() {
		err := srv.Start(fmt.Sprintf(":%s", port))
		log.Fatal(err)
	}()
}

func waitForShutdown(srv *server.Server) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			srv.Logger.Fatal(err)
		}
	}()
}