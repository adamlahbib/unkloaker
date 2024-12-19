package models

type Workload struct {
	ImageNameWithTag string
	Port             int32             `default:"8080"`
	Replicas         int32             `default:"1"`
	EnvVars          map[string]string `default:"{}"`
}

type CreateWorkloadRequest struct {
	Workload Workload
}

type CreateWorkloadResponse struct {
	Id      string
	Success bool
}

type DeleteWorkloadRequest struct {
	Id string
}

type DeleteWorkloadResponse struct {
	Success bool
}
