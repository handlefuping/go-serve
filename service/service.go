package service

import (
	"context"
	"fmt"
	"go-serve/registry"
	"net/http"
)

func Start(ctx context.Context, registration registry.Registration, registerHandler func()) context.Context {
	registerHandler()
	newCtx := startService(ctx, registration)
	registry.UnRegistry(registration)

	return newCtx

}

func startService (ctx context.Context, registration registry.Registration) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		err := http.ListenAndServe(registration.ServiceUrl, nil)
		if err != nil {
			fmt.Println(err)
		}
		cancel()
	}()
	go func ()  {
		fmt.Printf("%v is running Enter any key to stop\n", registration.ServiceName)
		var str string
		fmt.Scan(&str)
		cancel()
	}()

	return ctx
}