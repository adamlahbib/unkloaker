package models

type CloneRequest struct {
	Owner             string
	RepoName          string
	Url               string
	DockerfileContext string
}

type CloneResponse struct {
	Id                string
	Path              string
	DockerfileContext string
	Success           bool
}

type CleanupRequest struct {
	Id   string
	Path string
}

type CleanupResponse struct {
	Success bool
}
