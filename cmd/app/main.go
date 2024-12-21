package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/config"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/builder"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/pkg/database"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/pkg/server"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func main() {
	//load configuration via env
	cfg, err := config.NewConfig(".env")
	checkError(err)
	//init & start database
	db, err := database.InitDatabase(cfg.PostgresConfig)
	checkError(err)

	midtransClient := initMidtrans(cfg)

	//RBAC
	publicRoutes := builder.BuilderPublicRoutes(cfg, db, midtransClient)
	privateRoutes := builder.BuilderPrivateRoutes(cfg, db, midtransClient)
	//init & start server
	srv := server.NewServer(cfg, publicRoutes, privateRoutes)
	runServer(srv, cfg.PORT)
	waitForShutdown(srv)
}

func initMidtrans(cfg *config.Config) snap.Client {
	snapClient := snap.Client{}

	if cfg.ENV == "dev" {
		snapClient.New(cfg.MidtransConfig.ServerKey, midtrans.Sandbox)
	} else {
		snapClient.New(cfg.MidtransConfig.ServerKey, midtrans.Production)
	}

	return snapClient
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

func runServer(srv *server.Server, port string) {
	go func() {
		err := srv.Start(fmt.Sprintf(":%s", port))
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
