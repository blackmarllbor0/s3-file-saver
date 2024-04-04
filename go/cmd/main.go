package main

import (
	"app/internal/cfg"
	"app/internal/repository/postgres"
	"app/internal/repository/s3"
	"app/internal/transport/grpc"
	"app/pkg/runmode"
	"app/pkg/signal"
	"context"
	"log"
)

func main() {
	go signal.ListenSignals()

	runMode := runmode.GetAppRunMode()

	cfgService := cfg.NewConfigService()
	if err := cfgService.LoadConfig(runMode); err != nil {
		log.Fatalln(err)
	}

	_, err := s3.NewS3Session(cfgService)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = postgres.NewPgConnection(context.Background(), cfgService)

	gRPCConn, err := grpc.NewGRPCConnection(
		cfgService.GetTransportConfig().GRPC.Host,
		cfgService.GetTransportConfig().GRPC.Port,
	)
	if err != nil {
		log.Fatalln(err)
	}

	fileClient := grpc.NewFileWorkerClient(gRPCConn)
	_, err := fileClient.SaveFile(context.Background())
}
