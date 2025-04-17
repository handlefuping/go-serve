package service

import (
	"context"
	"fmt"
	"go-serve/registry"
	"net/http"
)

func Start(ctx context.Context, registration registry.RegistrationStruct, registerHandler func()) context.Context {
	registerHandler()
	newCtx := startService(ctx, registration)
	registry.RegistryService(registration)

	return newCtx

}

func startService (ctx context.Context, registration registry.RegistrationStruct) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		http.ListenAndServe(string(registration.ServiceUrl), nil)
		registry.UnRegistry(registration)
		
		cancel()
	}()
	go func ()  {
		fmt.Printf("%v is running Enter any key to stop\n", registration.ServiceName)
		var str string
		fmt.Scanln(&str)
		registry.UnRegistry(registration)
		cancel()
	}()

	return ctx
}