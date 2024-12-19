package servers

import (
	"context"
	"log"

	pb "github.com/adamlahbib/unkloaker/api/proto"
	"github.com/adamlahbib/unkloaker/models"
	"github.com/adamlahbib/unkloaker/services"
)

type KlonerServer struct {
	klonerService services.Repo
	pb.UnimplementedKlonerServer
}

func NewKloner(klonerService services.Repo) pb.KlonerServer {
	return &KlonerServer{klonerService: klonerService}
}

func (k *KlonerServer) Clone(ctx context.Context, req *pb.KlonerCloneRequest) (*pb.KlonerCloneResponse, error) {
	resp, err := k.klonerService.Clone(models.CloneRequest{
		Owner:             req.GetOwner(),
		RepoName:          req.GetReponame(),
		Url:               req.GetUrl(),
		DockerfileContext: req.GetDockerfilecontext(),
	})
	if err != nil {
		log.Printf("Error cloning repo: %v", err)
	}
	return &pb.KlonerCloneResponse{
		Id:                resp.Id,
		Path:              resp.Path,
		Dockerfilecontext: resp.DockerfileContext,
		Success:           resp.Success,
	}, err
}

func (k *KlonerServer) Cleanup(ctx context.Context, req *pb.KlonerCleanupRequest) (*pb.KlonerCleanupResponse, error) {
	resp, err := k.klonerService.Cleanup(models.CleanupRequest{
		Id: req.GetId(),
	})
	if err != nil {
		log.Printf("Error cleaning up repo: %v", err)
	}
	return &pb.KlonerCleanupResponse{
		Success: resp.Success,
	}, err
}
