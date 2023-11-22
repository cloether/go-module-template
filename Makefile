makefile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(notdir $(patsubst %/,%,$(dir $(makefile_path))))
envfile := .env

.PHONY: help build build-container rebuild-container start-container stop-container run logs test_coverage

# help target adapted from:
# https://gist.github.com/prwhite/8168133#gistcomment-2278355

TARGET_MAX_CHAR_NUM=20

## Start the services
start-container:
	@echo "Pulling images from Docker Hub"
	docker-compose pull
	@echo "Building Application Image"
	docker-compose build app
	@echo "Starting Docker Services"
	docker-compose up --detach
	./build/post-start.sh

build-container:
	@echo "building application image"
	docker-compose build app
	@echo "build completed. exit=$?"

rebuild-container:
	@echo "building application image (--no-cache)"
	docker-compose build --no-cache app
	@echo "build completed. exit=$?"

test_coverage:
	go test -coverprofile data/coverage "$1?" && go tool cover -html=data/coverage

## build the app
build:
	go build -i -v -ldflags="-X main.version=$(git describe --always --long --dirty)"

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

## Generate certificate private and public key
## Stop services
stop-container:
	docker-compose down
	docker volume rm $(current_dir)_{app,server}_node_modules 2>/dev/null || true

## Show env
$(envfile):
	@echo "Error: .env file does not exist! See the README for instructions."
	@exit 1
