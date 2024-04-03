package main

import (
	"app/internal/cfg"
	"app/internal/repository/postgres"
	"app/internal/repository/s3"
	"app/pkg/runmode"
	"app/pkg/signal"
	"context"
	"log"
	"os"
)

func main() {
	go signal.ListenSignals()

	runMode := runmode.GetAppRunMode()

	cfgService := cfg.NewConfigService()
	if err := cfgService.LoadConfig(runMode); err != nil {
		log.Fatalln(err)
	}

	s3Repo, err := s3.NewS3Session(cfgService)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = postgres.NewPgConnection(context.Background(), cfgService)

	file, err := os.Open("config.dev.yaml")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}

	if err := s3Repo.SaveFile(file); err != nil {
		log.Fatalln(err)
	}

	if err := s3Repo.DeleteFile(file.Name()); err != nil {
		log.Fatalln(err)
	}

	log.Println("Good test")
}
