package registry

import "sync"



type providerStruct struct {
	services map[ServiceName][]ServiceUrl
	mutex *sync.Mutex
}

var provider = providerStruct{
	services: make(map[ServiceName][]ServiceUrl),
	mutex: new(sync.Mutex),
}

func (provider *providerStruct) update(patch patchStruct) {
	provider.mutex.Lock()
	defer provider.mutex.Unlock()

	for _, patchEntry := range patch.added {
		if services, ok := provider.services[patchEntry.serviceName]; ok {
			provider.services[patchEntry.serviceName] = append(services,patchEntry.serviceUrl)
		} else {
			provider.services[patchEntry.serviceName] = []ServiceUrl{patchEntry.serviceUrl}
		}
	}
	
	for _, patchEntry := range patch.removed {
		if services, ok := provider.services[patchEntry.serviceName]; ok {
			for i, url := range services {
				if url == patchEntry.serviceUrl {
					provider.services[patchEntry.serviceName] = append(services[:i], services[i+1:]...)
				}
			}
		}
	}

}
func (provider *providerStruct) Get(name ServiceName) ServiceUrl {
	if _, ok := provider.services[name]; ok {
		return provider.services[name][0]
	}

	return ""
}