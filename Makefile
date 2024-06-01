#!make
include .make.env

DOCKER_COMPOSE ?= docker-compose
DOCKER_COMPOSE_FILE ?= docker-compose.yaml
# colors
CYAN=\033[0;36m
NC=\033[0m

.PHONY: help

## Show help screen
help:
	@echo "[~~~~~ $(CYAN)HELP$(NC) ~~~~~]"
	@printf "Available targets:\n\n"
	@awk '/^[a-zA-Z\-\_0-9%:\\]+/ { \
	  helpMessage = match(lastLine, /^## (.*)/); \
	  if (helpMessage) { \
		helpCommand = $$1; \
		helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
  gsub("\\\\", "", helpCommand); \
  gsub(":+$$", "", helpCommand); \
		printf "  \x1b[32;01m%-35s\x1b[0m %s\n", helpCommand, helpMessage; \
	  } \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST) | sort -u
	@printf "\n"

## Initialize project
init:
	@echo "[~~~~~ $(CYAN)INIT$(NC) ~~~~~]"
	@echo "installing golangci-lint 🧹"
	@make install/golangci-lint
	@echo "installing mockgen 🧹"
	@make install/mockgen
	@echo "installing go test 🧹"
	@go install github.com/rakyll/gotest@latest
	@echo "Done ✅"

## Update dependencies
dep:
	@echo "[~~~~~ $(CYAN)DEP$(NC) ~~~~~]"
	@echo "cleaning module cache 🧹"
	@go clean -modcache
	@echo "tidying up "
	@go mod tidy
	@echo "update vendor ⏫"
	@go mod vendor
	@echo "Done ✅"

## Running docker compose for targeted service'
dev/up:
	@echo "[~~~~~ $(CYAN)DEV$(NC) ~~~~~]"
	@echo "running service $(SERVICE) using $(DOCKER_COMPOSE) 💨"
	@$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) up --build $(SERVICE)

## Remove docker container by compose config
dev/down:
	@echo "[~~~~~ $(CYAN)DEV CLEAR$(NC) ~~~~~]"
	@echo "Removing docker container service 🧹"
	@$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) down

install/db_migration:
	@echo "[~~~~~ $(CYAN)INSTALL DB MIGRATOR$(NC) ~~~~~]"
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@$(DB_MIGRATOR_TAG)

install/mockgen:
	@echo "[~~~~~ $(CYAN)INSTALL MOCKGEN$(NC) ~~~~~]"
	@go install go.uber.org/mock/mockgen@latest

install/golangci-lint:
	@echo "[~~~~~ $(CYAN)INSTALL GOLANGCI-LINT$(NC) ~~~~~]"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53
