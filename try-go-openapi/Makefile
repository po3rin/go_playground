validate:
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v3.3.4 validate -i /local/schema.yml
generate-go:
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v3.3.4 generate \
    --skip-validate-spec \
    -i /local/schema.yml \
    -g go-gin-server \
    -o /local/server/
generate-js:
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v3.3.4 generate \
    --skip-validate-spec \
    -i /local/schema.yml \
    -g javascript \
    -o /local/client/
