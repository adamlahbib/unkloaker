package servers

import (
	"context"
	"log"

	pb "github.com/adamlahbib/unkloaker/api/proto"
	"github.com/adamlahbib/unkloaker/models"
	"github.com/adamlahbib/unkloaker/services"
)

type UnkloakerServer struct {
	unkloakerService services.Workload
	pb.UnimplementedUnkloakerServer
}

func NewUnkloaker(unkloakerService services.Workload) pb.UnkloakerServer {
	return &UnkloakerServer{unkloakerService: unkloakerService}
}

func (u *UnkloakerServer) Deploy(ctx context.Context, req *pb.UnkloakerDeployRequest) (*pb.UnkloakerDeployResponse, error) {
	resp, err := u.unkloakerService.Create(ctx, models.CreateWorkloadRequest{
		Workload: models.Workload{
			ImageNameWithTag: req.GetImagenamewithtag(),
			Port:             req.GetPort(),
			Replicas:         req.GetReplicas(),
			EnvVars:          req.GetEnvvars(),
		},
	})
	if err != nil {
		log.Printf("Error deploying workload: %v", err)
	}
	return &pb.UnkloakerDeployResponse{
		Id:      resp.Id,
		Success: resp.Success,
	}, err
}

func (u *UnkloakerServer) Undeploy(ctx context.Context, req *pb.UnkloakerUndeployRequest) (*pb.UnkloakerUndeployResponse, error) {
	resp, err := u.unkloakerService.Delete(ctx, models.DeleteWorkloadRequest{
		Id: req.GetId(),
	})
	if err != nil {
		log.Printf("Error undeploying workload: %v", err)
	}
	return &pb.UnkloakerUndeployResponse{
		Success: resp.Success,
	}, err
}
