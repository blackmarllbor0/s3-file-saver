package main

import (
	"app/internal/cfg"
	"app/internal/repository/postgres"
	"app/internal/repository/s3"
	"app/internal/service"
	"app/internal/transport/grpcClient"
	"app/pkg/logger"
	"app/pkg/runmode"
	"app/pkg/signal"
	"context"
	"log"
)

func main() {
	go signal.ListenSignals()

	runMode := runmode.GetAppRunMode()

	logrusLogger, err := logger.NewLogrusLogger(runMode)
	if err != nil {
		log.Fatalln(err)
	}

	cfgService := cfg.NewConfigService(logrusLogger)
	if err := cfgService.LoadConfig(runMode); err != nil {
		log.Fatalln(err)
	}

	_, err = postgres.NewPgConnection(context.Background(), cfgService, logrusLogger)

	fileMetadataRepo := postgres.NewFileMetadataRepo()
	fileRepo, err := s3.NewS3Session(cfgService, logrusLogger)
	if err != nil {
		log.Fatalln(err)
	}

	fileWorkerService := service.NewFileWorkerService(fileRepo, fileMetadataRepo)

	server := grpcClient.NewGRPCServer(fileWorkerService, cfgService, logrusLogger)
	if err := server.ListenGRPCServer(); err != nil {
		log.Fatalln(err)
	}

	defer server.Stop()
}
