package registry

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)


func RegistryService(registration RegistrationStruct) {
	b, err := json.Marshal(registration)
	if err != nil {
		log.Println("stringify registration error")
		return
	}
	// http.Handle(string(registration.ServiceUpdateUrl), &updateHandlerStruct{})
	res, err := http.Post(ServerAddress, "application/json", bytes.NewReader(b))
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

func UnRegistry(registration RegistrationStruct) {
	req, err := http.NewRequest(http.MethodDelete, ServerAddress, bytes.NewReader([]byte(registration.ServiceUrl)))
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


type updateHandlerStruct struct {}
func (updateHandler updateHandlerStruct)  ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	decoder := json.NewDecoder(r.Body)
	var patch patchStruct
	err := decoder.Decode(&patch)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	provider.update(patch)
}

