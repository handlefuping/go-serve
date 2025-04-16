package main

import (
	"context"
	"go-serve/log"
	"go-serve/registry"
	"go-serve/service"
	stLog "log"
)


func main() {
	log.Run("./destination.log")

	ctx := service.Start(context.Background(), registry.Registration{
		ServiceName: "log service",
		ServiceUrl: "localhost:8000",
	}, log.LogHandler)

	<- ctx.Done()
	stLog.Fatalln("log service shut down")
}