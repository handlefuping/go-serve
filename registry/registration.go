package registry


type ServiceName string
type ServiceUrl string

type RegistrationStruct struct {
	ServiceName ServiceName
	ServiceUrl ServiceUrl
	RequiredServices []ServiceName
	ServiceUpdateUrl ServiceUrl
}

const (
	LogService     = ServiceName("LogService")
)

type patchEntryStruct struct {
	serviceName ServiceName
	serviceUrl ServiceUrl
}

type patchStruct struct {
	added []patchEntryStruct
	removed []patchEntryStruct
}
