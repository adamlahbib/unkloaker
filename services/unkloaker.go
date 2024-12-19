package services

import (
	"context"

	"github.com/adamlahbib/unkloaker/helpers"
	"github.com/adamlahbib/unkloaker/models"
	"github.com/google/uuid"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Workload interface {
	Create(ctx context.Context, request models.CreateWorkloadRequest) (models.CreateWorkloadResponse, error)
	Delete(ctx context.Context, request models.DeleteWorkloadRequest) (models.DeleteWorkloadResponse, error)
}

type Unkloaker struct {
	kubeClient *kubernetes.Clientset
}

func NewUnkloaker(client *kubernetes.Clientset) Workload {
	return &Unkloaker{kubeClient: client}
}

func (u *Unkloaker) Create(ctx context.Context, request models.CreateWorkloadRequest) (models.CreateWorkloadResponse, error) {
	id := uuid.New().String()
	deployment := helpers.PrepareDeployment(id, request.Workload.ImageNameWithTag, request.Workload.Port, request.Workload.Replicas, request.Workload.EnvVars)
	service := helpers.PrepareService(id, request.Workload.Port)
	if _, err := u.kubeClient.AppsV1().Deployments("unkloak").Create(ctx, deployment, metav1.CreateOptions{}); err != nil {
		return models.CreateWorkloadResponse{Id: id, Success: false}, err
	}
	if _, err := u.kubeClient.CoreV1().Services("unkloak").Create(ctx, service, metav1.CreateOptions{}); err != nil {
		return models.CreateWorkloadResponse{Id: id, Success: false}, err
	}
	return models.CreateWorkloadResponse{Id: id, Success: true}, nil
}

func (u *Unkloaker) Delete(ctx context.Context, request models.DeleteWorkloadRequest) (models.DeleteWorkloadResponse, error) {
	deletePolicy := metav1.DeletePropagationForeground // delete pods before the deployment
	if err := u.kubeClient.AppsV1().Deployments("unkloak").Delete(ctx, request.Id, metav1.DeleteOptions{PropagationPolicy: &deletePolicy}); err != nil {
		return models.DeleteWorkloadResponse{Success: false}, err
	}
	if err := u.kubeClient.CoreV1().Services("unkloak").Delete(ctx, request.Id, metav1.DeleteOptions{}); err != nil {
		return models.DeleteWorkloadResponse{Success: false}, err
	}
	return models.DeleteWorkloadResponse{Success: true}, nil
}
