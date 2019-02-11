PHONY: validate generate-go generate-js help

validate: ## validate schema.yml
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v3.3.4 validate -i /local/schema.yml

generate-go: ## generate go server
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v3.3.4 generate \
    -i /local/schema.yml \
    -g go-gin-server \
    -o /local/server/

generate-js: ## generate js client
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v3.3.4 generate \
    -i /local/schema.yml \
    -g javascript \
    -o /local/client/openapi/

mock: ## run mock server
	docker run -p 8080:8080 -v ${PWD}:/local danielgtaylor/apisprout /local/schema.yml -p 8080

help: ## Display help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
