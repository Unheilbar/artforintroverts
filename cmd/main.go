package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	adapter "github.com/unheilbar/artforintrovert_entry_task/internal/adapters"
	"github.com/unheilbar/artforintrovert_entry_task/internal/api"
	"github.com/unheilbar/artforintrovert_entry_task/internal/config"
	"github.com/unheilbar/artforintrovert_entry_task/internal/scheduler"
	"github.com/unheilbar/artforintrovert_entry_task/internal/server"
	"github.com/unheilbar/artforintrovert_entry_task/internal/service/user"
	"github.com/unheilbar/artforintrovert_entry_task/internal/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO
// config
// structure
// bd connection

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	opts := options.Client().ApplyURI(config.DSN)
	mongoClient, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}

	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	provider := storage.NewStorageProvider(mongoClient, config.DbName, config.CollectionName)

	ctx, cancel := context.WithCancel(context.Background())

	go scheduler.RunRefreshJob(ctx, provider.RefreshCache, config.CacheRefreshInterval)

	storageProvider := adapter.NewStorageProviderAdapter(provider)

	service := user.NewService(storageProvider)

	handlers := api.NewHandlers(service)

	server := server.NewServer(handlers.GetRouter(), config.ServerAddress)

	serverErrChan := server.Start()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-signals:
		fmt.Println("main: got terminate signal. Shutting down...", nil)
		cancel()
	case <-serverErrChan:
		fmt.Println("main: got server err signal. Shutting down...", nil)
	}
}
