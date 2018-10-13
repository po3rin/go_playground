.PHONY: mock
mock: ## run mock server.
	go run mock/main.go

.PHONY: server
server: ## run api server.
	go run main.go

# Help =====================================

help: ## Display help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
