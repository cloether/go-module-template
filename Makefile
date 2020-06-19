mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))
envfile := ./.moolah.json  # TODO: env file is now yaml. Change to handle yaml files

.PHONY: help build rebuild cert start stop run logs local config coverage webhook graph_deps

# help target adapted from: https://gist.github.com/prwhite/8168133#gistcomment-2278355

TARGET_MAX_CHAR_NUM=20

## Start the services
start:
	@echo "Pulling images from Docker Hub"
	docker-compose pull
	@echo "Building Application Image"
	docker-compose build app
	@echo "Starting Docker Services"
	docker-compose up --detach
	./build/post-start.sh
build:
	@echo "building application image"
	docker-compose build app
	@echo "build completed. exit=$?"
rebuild:
	@echo "building application image (--no-cache)"
	docker-compose build --no-cache app
	@echo "build completed. exit=$?"
coverage:
	go test -coverprofile data/coverage "$1?" && go tool cover -html=data/coverage

## Start the services locally
run:
	go run main.go

## Show help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  make <target>'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z_0-9-]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  %-$(TARGET_MAX_CHAR_NUM)s %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

## Show the service logs (services must be running)
logs:
	docker-compose logs --follow

## Generate certificate priavte and public key
## Stop services
stop:
	docker-compose down
	docker volume rm $(current_dir)_{app,server}_node_modules 2>/dev/null || true

## Show env
$(envfile):
	@echo "Error: .env file does not exist! See the README for instructions."
	@exit 1
