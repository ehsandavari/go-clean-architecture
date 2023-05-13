package constants

type ServiceMode string

func (r ServiceMode) IsValid() bool {
	if r == ServiceModeDevelopment || r == ServiceModeStage || r == ServiceModeProduction {
		return true
	}
	return false
}

func (r ServiceMode) String() string {
	return string(r)
}

const (
	ServiceModeDevelopment ServiceMode = "development"
	ServiceModeStage       ServiceMode = "stage"
	ServiceModeProduction  ServiceMode = "production"
)
