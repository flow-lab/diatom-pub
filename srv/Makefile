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
generate g: api-gen-models ## Run all code generation
	sqlc generate
	go generate -run="mockgen" ./...

.PHONY: api-gen-models agm
api-gen-models: ## Generate models from api
	./gen-models.sh

.PHONY: test t
test t: ## Run tests
	go test -mod=readonly -covermode=atomic -v ./...

.PHONY: dev-watch dw
dev-watch dw: ## Run the server in dev mode with watch gow, install https://github.com/mitranim/gow
	# add all files that need to be compiled
	env PORT=8888 DB_HOST=localhost DB_PORT=5432 DB_NAME=diatom-dev DB_USERNAME=diatom-dev DB_PASSWORD=diatom-dev REDIS_HOST=localhost REDIS_PORT=6379 gow run cmd/srv/main.go cmd/srv/health.go cmd/srv/apidoc.go srv/db/.

.PHONY: api-lint al
api-lint: g ## Lint api using
	npx @redocly/cli lint .api.yaml

.PHONY: api-preview ap
api-preview: g ## Lint api using
	npx @redocly/cli preview-docs .api.yaml

.PHONY: help
help: ## Show this
	@grep -E '^[a-zA-Z_ -]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-21s\033[0m %s\n", $$1, $$2}'