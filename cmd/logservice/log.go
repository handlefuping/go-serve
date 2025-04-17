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
	serviceAddress := "http://localhost:8080"
	ctx := service.Start(context.Background(), registry.RegistrationStruct{
		ServiceName: registry.LogService,
		ServiceUrl: registry.ServiceUrl(serviceAddress),
		RequiredServices: make([]registry.ServiceName, 0),
		ServiceUpdateUrl: registry.ServiceUrl(serviceAddress) + "/services",
	}, log.LogHandler)

	<- ctx.Done()
	stLog.Fatalln("log service shut down")
}