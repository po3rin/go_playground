validate:
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli validate -i /local/schema.yml
generate:
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
    -i https://raw.githubusercontent.com/openapitools/openapi-generator/master/modules/openapi-generator/src/test/resources/2_0/petstore.yaml \
    -g go-server \
    -o /local/server/
