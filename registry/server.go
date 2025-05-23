package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

const (
	ServerPort = ":8000"
	ServerAddress = "http://localhost" + ServerPort +  "/services"
)

type RegistryStruct struct {
	registrations []RegistrationStruct
	mutex *sync.Mutex
}

func (registry *RegistryStruct) add(registration RegistrationStruct) {
	registry.mutex.Lock()
	registry.registrations = append(registry.registrations, registration)
	registry.mutex.Unlock()
	registry.sendRequireServices(registration)
}

func (registry *RegistryStruct) sendRequireServices (registration RegistrationStruct) {
	registry.mutex.Lock()
	var patch patchStruct
	for _, reg := range registry.registrations {
		for _, serviceName := range registration.RequiredServices {
			if serviceName == reg.ServiceName {
				patch.added = append(patch.added, patchEntryStruct{serviceName: serviceName, serviceUrl: reg.ServiceUrl})
			}
		}
	}
	registry.mutex.Unlock()
	registry.sendPatch(patch, registration.ServiceUpdateUrl)

} 
func (registry *RegistryStruct) sendPatch (patch patchStruct, url ServiceUrl) {
	bt, err:= json.Marshal(patch)
	if err != nil {
		fmt.Println(err)
	} else {
		http.Post(string(url), "application/json", bytes.NewReader(bt))
	}
} 
func (registry *RegistryStruct) remove(url ServiceUrl) {
	for index, registration := range registry.registrations {
		if registration.ServiceUrl == url {
			registry.mutex.Lock()
			registry.registrations = append(registry.registrations[:index], registry.registrations[index+1:]...)
			registry.mutex.Unlock()
		}
	}
}

var registry *RegistryStruct

func Run() {
	registry = &RegistryStruct{
		mutex: new(sync.Mutex),
		registrations: make([]RegistrationStruct, 1),
	}
}

type RegistryHandlerStruct struct{} 

func (registryHandler *RegistryHandlerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodPost: 
			msg, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return 
			}
			var registration RegistrationStruct
			err = json.Unmarshal(msg, &registration)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return 
			}
			registry.add(registration)
			w.WriteHeader(http.StatusOK)
		case http.MethodDelete:
			url, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return 
			}
			registry.remove(ServiceUrl(url))
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
}



