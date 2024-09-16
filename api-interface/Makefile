project_name = api-interface
image_name = gofiber:latest

# Charger les variables d'environnement depuis le fichier .env
ifneq (,$(wildcard ../.env))
    include ../.env
    export $(shell ./load_env.exe)
endif

help: ## This help dialog.
	@grep -F -h "##" $(MAKEFILE_LIST) | grep -F -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

run-local: ## Run the app locally
	go run app.go

requirements: ## Generate go.mod & go.sum files
	go mod tidy

clean-packages: ## Clean packages
	go clean -modcache

up: ## Run the project in a local container
	make up-silent
	make shell

build: ## Generate docker image
	docker build -t $(image_name) .

build-no-cache: ## Generate docker image with no cache
	docker build --no-cache -t $(image_name) .

up-silent: ## Run local container in background
	make delete-container-if-exist
	docker run -d -p $(S3_PORT):$(S3_PORT) --name $(project_name) $(image_name) ./app

up-dev: ## Run local container in foreground to see logs
	make delete-container-if-exist
	docker run -p $(S3_PORT):$(S3_PORT) --name $(project_name) -v ${PWD}:/app $(image_name) ./app

up-silent-prefork: ## Run local container in background with prefork
	make delete-container-if-exist
	docker run -d -p $(S3_PORT):$(S3_PORT) --name $(project_name) $(image_name) ./app -prod

delete-container-if-exist: ## Delete container if it exists
	docker stop $(project_name) || exit 0 && docker rm $(project_name) || exit 0

shell: ## Run interactive shell in the container
	docker exec -it $(project_name) /bin/sh

stop: ## Stop the container
	docker stop $(project_name)

start: ## Start the container
	docker start $(project_name)

start-dev: ## Start the project in development mode with file watching
	air