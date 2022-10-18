SHELL := /bin/bash -e

.PHONY: docker-up du
docker-up du: ## docker up and build
	sudo docker compose --profile prod up --build -d

.PHONY: docker-up-recreate dur
docker-up-recreate dur: ## docker up force recreate and build
	sudo docker compose --profile prod up --force-recreate --build -d

.PHONY: help
help: ## Show this
	@grep -E '^[a-zA-Z_ -]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-22s\033[0m %s\n", $$1, $$2}'