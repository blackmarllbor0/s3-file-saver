package grpc

import (
	"app/internal/cfg"
	"app/internal/service"
	"fmt"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

type GRPCServer struct {
	cfgService cfg.ConfigService

	fileWorkerService service.FileWorkerService
}

func NewGRPCServer(cfgService cfg.ConfigService, fileWorkerService service.FileWorkerService) *GRPCServer {
	return &GRPCServer{
		cfgService:        cfgService,
		fileWorkerService: fileWorkerService,
	}
}

func (g GRPCServer) ListenGRPCServer() error {
	addr := fmt.Sprintf("%s:%d", g.cfgService.GetTransportConfig().GRPC.Host, g.cfgService.GetTransportConfig().GRPC.Port)
	liestener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("grpc: failed to listen tcp: %v", err)
	}

	server := grpc.NewServer()

	RegisterFileWorkerServer(s, g.fileWorkerService)

	if err := server.Serve(liestener); err != nil {
		return fmt.Errorf("grpc: failed to serve server: %v", err)
	}

	log.Println("")

	return nil
}
