package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


func Registry(r Registration) {
	b, err := json.Marshal(r)
	if err != nil {
		log.Println("stringify registration error")
		return
	}
	res, err := http.Post(fmt.Sprintf("http://%v/services", r.ServiceUrl), "application/json", bytes.NewReader(b))
	if err != nil {
		log.Println("registry service error")
	}
	if res.StatusCode != http.StatusOK {
		log.Println("registry service error")
	}
	if err != nil || res.StatusCode != http.StatusOK{
		log.Println("registry service response status is not ok", res.StatusCode)
	}
}

func UnRegistry(r Registration) {
	fmt.Println(fmt.Sprintf("http://%v/services", r.ServiceUrl))
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%v/services", r.ServiceUrl), bytes.NewReader([]byte(r.ServiceUrl)))
	if err != nil {
		log.Println("unregistry service error")
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("do unregistry service error")
	}
	if res.StatusCode != http.StatusOK{
		log.Println("unregistry service response status is not ok", res.StatusCode)
	}

}
