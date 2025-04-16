package registry

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)




type Registration struct {
	ServiceName string
	ServiceUrl string
}

type pool struct {
	collection []Registration
	mu *sync.Mutex
}

func (p *pool) add(r Registration) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.collection = append(p.collection, r)
}

func (p *pool) remove(url string) {
	for i, v := range p.collection {
		if v.ServiceUrl == url {
			p.mu.Lock()
			p.collection = append(p.collection[:i], p.collection[i+1:]...)
			p.mu.Unlock()
		}
	}
}
var rgPool *pool

func Run() {
	rgPool = &pool{
		mu: new(sync.Mutex),
		collection: make([]Registration, 10),
	}
}

type RegistryHandler struct{} 

func (rh *RegistryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodPost: 
			msg, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return 
			}
			var r Registration
			err = json.Unmarshal(msg, &r)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return 
			}
			rgPool.add(r)
			w.WriteHeader(http.StatusOK)
		case http.MethodDelete:
			fmt.Println(11222)
			url, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return 
			}
			rgPool.remove(string(url))
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
}
