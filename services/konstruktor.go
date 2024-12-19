package services

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/adamlahbib/unkloaker/models"
	"github.com/docker/docker/api/types"
	docker "github.com/docker/docker/client"
)

type Image interface {
	Build(ctx context.Context, request models.ImageRequest) (models.ImageResponse, error)
}

type Konstruktor struct {
	dockerClient *docker.Client
}

func NewKonstruktor(client *docker.Client) Image {
	return &Konstruktor{dockerClient: client}
}

func (k *Konstruktor) Build(ctx context.Context, request models.ImageRequest) (models.ImageResponse, error) {
	// open the dockerfile
	dockerfile, err := os.Open(filepath.Join(request.Image.RepoPath, request.Image.DockerfileContext, "Dockerfile"))
	if err != nil {
		return models.ImageResponse{ImageNameWithTag: "", Success: false}, err
	}
	defer dockerfile.Close()

	// create a tar archive from the dockerfile
	tar, err := os.CreateTemp("", "dockerfile.tar")
	if err != nil {
		return models.ImageResponse{ImageNameWithTag: "", Success: false}, err
	}
	defer tar.Close()

	// build the docker image
	imageNameWithTag := request.Image.Name + ":" + request.Image.Tag
	imageBuildResponse, err := k.dockerClient.ImageBuild(ctx, tar, types.ImageBuildOptions{
		Dockerfile: "Dockerfile",
		Context:    tar,
		Tags:       []string{imageNameWithTag},
	})
	if err != nil {
		return models.ImageResponse{ImageNameWithTag: "", Success: false}, err
	}
	defer imageBuildResponse.Body.Close()

	// output the build response
	if _, err := io.Copy(os.Stdout, imageBuildResponse.Body); err != nil {
		panic(err)
	}

	return models.ImageResponse{ImageNameWithTag: imageNameWithTag, Success: true}, nil
}
