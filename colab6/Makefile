PHONY: validate generate-go stub

validate: ## validate schema.yml
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v3.3.4 validate -i /local/schema/schema.yml

generate-go: ## generate go server
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v3.3.4 generate \
    -i /local/schema/schema.yml \
    -g go-gin-server \
    -o /local/server/

stub: ## run mock server
	docker run -p 8080:8080 \
      -v ${PWD}:/local danielgtaylor/apisprout \
      /local/schema/schema.yml -p 8080

editor: ## run editor
	docker run -p 3000:8080 swaggerapi/swagger-editor