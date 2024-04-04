RUN_MODE ?= dev

include .env.$(RUN_MODE)

.PHONY:build
build:
	docker-compose --env-file .env.$(RUN_MODE) up -d

.PHONY:down
down:
	docker-compose down

.PHONY:migrate-up
migrate-up:
	migrate -database $(POSTGRES_CONN_STRING) -path ./migrations/pg up

.PHONY:migrate-down
migrate-down:
	migrate -database $(POSTGRES_CONN_STRING) -path ./migrations/pg down