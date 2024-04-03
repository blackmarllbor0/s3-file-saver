package main

import (
	"app/internal/cfg"
	"app/internal/repository/s3"
	"app/pkg/runmode"
	"app/pkg/signal"
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

	file, err := os.Open("cmd/main.go")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}

	if err := s3Repo.SaveFile(file); err != nil {
		log.Fatalln(err)
	}
}
