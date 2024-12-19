package services

import (
	"os"

	"github.com/adamlahbib/unkloaker/models"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/google/uuid"
)

type Repo interface {
	Clone(request models.CloneRequest) (models.CloneResponse, error)
	Cleanup(request models.CleanupRequest) (models.CleanupResponse, error)
}

type Kloner struct {
	Repo Repo
}

func NewKloner() Repo {
	return &Kloner{}
}

var Token = os.Getenv("GITHUB_TOKEN")

func (k *Kloner) Clone(request models.CloneRequest) (models.CloneResponse, error) {
	if _, err := git.PlainClone(
		"/tmp/"+request.RepoName, false, &git.CloneOptions{
			URL:  request.Url,
			Auth: &http.BasicAuth{Username: "token", Password: Token},
		}); err != nil {
		return models.CloneResponse{
			Id:                uuid.New().String(),
			Path:              "/tmp/" + request.RepoName,
			DockerfileContext: request.DockerfileContext,
			Success:           false}, err
	}
	return models.CloneResponse{
		Id:                uuid.New().String(),
		Path:              "/tmp/" + request.RepoName,
		DockerfileContext: request.DockerfileContext,
		Success:           true}, nil
}

func (k *Kloner) Cleanup(request models.CleanupRequest) (models.CleanupResponse, error) {
	if err := os.RemoveAll(request.Path); err != nil {
		return models.CleanupResponse{Success: false}, err
	}
	return models.CleanupResponse{Success: true}, nil
}
