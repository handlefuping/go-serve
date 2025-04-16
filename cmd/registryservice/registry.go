package main

import (
	"fmt"
	"go-serve/registry"
	"log"
	"net/http"
)


func main() {
	registry.Run()
	http.Handle("/services", &registry.RegistryHandler{})
	ch := make(chan bool)
	go func() {
		fmt.Println("registry service is running")
		log.Fatal(http.ListenAndServe(":8000", nil))
		ch <- true
	}()

	go func ()  {
		fmt.Println("registry service is running Enter any key to stop")
		var str string
		fmt.Scan(&str)
		ch <- true
	}()
	<- ch

	fmt.Println("registry service shut down")
}