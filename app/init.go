package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	pb "github.com/adamlahbib/unkloaker/api/proto"
	"github.com/adamlahbib/unkloaker/api/servers"
	"github.com/adamlahbib/unkloaker/services"
	"github.com/docker/docker/client"
	"google.golang.org/grpc"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func Start() {
	// prepare kubernetes and docker clients as well as the services
	kubeConfigPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	konfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		log.Fatalf("Error loading kubeconfig: %s", err)
	}
	klient, err := kubernetes.NewForConfig(konfig)
	if err != nil {
		log.Fatalf("Error creating kubernetes client: %s", err)
	}
	unkloakerService := servers.NewUnkloaker(services.NewUnkloaker(klient))

	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Error creating docker client: %s", err)
	}
	konstruktorService := servers.NewKonstruktor(services.NewKonstruktor(dockerClient))
	klonerService := servers.NewKloner(services.NewKloner())

	// grpc service initialization
	port := 5000
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("Error creating listener: %s", err)
	}
	log.Println("starting gRPC server on port", port)
	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	pb.RegisterUnkloakerServer(grpcServer, unkloakerService)
	pb.RegisterKonstruktorServer(grpcServer, konstruktorService)
	pb.RegisterKlonerServer(grpcServer, klonerService)
	sCtx := serverContext(context.Background())
	go func() {
		log.Printf("gRPC server listening on port: %d", listener.Addr())
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	<-sCtx.Done()
	grpcServer.GracefulStop()
	log.Println("gRPC server stopped")
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		s := <-c
		log.Printf("received signal: %v", s)
		cancel()
	}()
	return ctx
}
