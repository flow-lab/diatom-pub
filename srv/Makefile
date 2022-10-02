SHELL := /bin/bash -e

# ----------------------------------------------------------------------------------------------------------------------
# dev
# ----------------------------------------------------------------------------------------------------------------------

.PHONY: deps d
deps d: ## Download dependencies
	go mod download

.PHONY: deps-upgrade du
deps-upgrade du: ## Upgrade all dependencies
	go get -u -t ./...

.PHONY: deps-reset dr
deps-reset dr: ## Reset changes in go.mod
	git checkout -- go.mod

.PHONY: tidy ti
tidy ti: ## Tidy the dependencies
	go mod tidy

.PHONY: verify ve
verify ve: ## Verify dependencies
	go mod verify

.PHONY: generate g
generate g: ## Run all code generation
	sqlc generate
	go generate -run="mockgen" ./...

.PHONY: test t
test t: ## Run tests
	go test -mod=readonly -covermode=atomic -v ./...

.PHONY: dev-watch dw
dev-watch dw: ## Run the server in dev mode with watch gow, install https://github.com/mitranim/gow
	# add all files that need to be compiled
	env PORT=8888 DB_HOST=localhost DB_PORT=5432 DB_NAME=diatom-dev DB_USERNAME=diatom-dev DB_PASSWORD=diatom-dev REDIS_HOST=localhost REDIS_PORT=6379 gow run cmd/srv/main.go cmd/srv/health.go cmd/srv/apidoc.go srv/db/.

api-lint:
	curl http://localhost/backend/api/api.yaml > api.yaml
	openapi lint api.yaml

.PHONY: help
help: ## Show this
	@grep -E '^[a-zA-Z_ -]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-21s\033[0m %s\n", $$1, $$2}'