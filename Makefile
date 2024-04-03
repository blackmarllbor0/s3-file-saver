RUN_MODE ?= dev

include .env.$(RUN_MODE)

.PHONY:build
build:
	docker-compose --env-file .env.$(RUN_MODE) up -d

.PHONY:down
down:
	docker-compose down
