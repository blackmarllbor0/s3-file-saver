package grpcClient

import (
	"app/internal/cfg"
	"app/pkg/logger"
	"fmt"
	"net"

	grpc "google.golang.org/grpc"
)

type GRPCServer struct {
	server *grpc.Server

	fileWorkerService FileWorkerServer
	cfgService        cfg.ConfigService
	logService        logger.LoggerService
}

func NewGRPCServer(
	fileWorkerService FileWorkerServer,
	cfgService cfg.ConfigService,
	logService logger.LoggerService,
) *GRPCServer {
	return &GRPCServer{
		cfgService:        cfgService,
		fileWorkerService: fileWorkerService,
		logService:        logService,
	}
}

func (g *GRPCServer) ListenGRPCServer() error {
	addr := fmt.Sprintf("%s:%d", g.cfgService.GetTransportConfig().GRPC.Host, g.cfgService.GetTransportConfig().GRPC.Port)
	liestener, err := net.Listen("tcp", addr)
	if err != nil {
		err := fmt.Errorf("grpcClient.GRPCServer.ListenGRPCServer: failed to listen tcp: %v", err)

		g.logService.Error(err.Error())

		return err
	}

	g.server = grpc.NewServer()

	RegisterFileWorkerServer(g.server, g.fileWorkerService)
	go g.logService.Info(
		"grpcClient.GRPCServer.ListenGRPCServer: gr" +
			"pc services have been successfully registered",
	)

	go g.logService.Info(fmt.Sprintf("grpcClient.GRPCServer.ListenGRPCServer: grpc server listen on %s", addr))

	if err := g.server.Serve(liestener); err != nil {
		err := fmt.Errorf("grpcClient.GRPCServer.ListenGRPCServer: failed to serve server: %v", err)

		g.logService.Error(err.Error())

		return err
	}

	return nil
}

func (g *GRPCServer) Stop() {
	g.server.Stop()
	g.logService.Info("grpcClient.GRPCServer.Stop: grpc server successfully stopped")
}
