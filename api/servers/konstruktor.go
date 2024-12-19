package servers

import (
	"context"
	"log"

	pb "github.com/adamlahbib/unkloaker/api/proto"
	"github.com/adamlahbib/unkloaker/models"
	"github.com/adamlahbib/unkloaker/services"
)

type KonstruktorServer struct {
	konstruktorService services.Image
	pb.UnimplementedKonstruktorServer
}

func NewKonstruktor(konstruktorService services.Image) pb.KonstruktorServer {
	return &KonstruktorServer{konstruktorService: konstruktorService}
}

func (k *KonstruktorServer) Build(ctx context.Context, req *pb.KonstruktorBuildRequest) (*pb.KonstruktorBuildResponse, error) {
	resp, err := k.konstruktorService.Build(ctx, models.ImageRequest{
		Image: models.Image{
			RepoPath:          req.GetRepopath(),
			DockerfileContext: req.GetDockerfilecontext(),
			Name:              req.GetName(),
			Tag:               req.GetTag(),
		},
	})
	if err != nil {
		log.Println("Error building image: ", err)
	}
	return &pb.KonstruktorBuildResponse{
		Imagenamewithtag: resp.ImageNameWithTag,
		Success:          resp.Success,
	}, err
}
