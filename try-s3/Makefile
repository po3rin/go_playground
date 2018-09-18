IMAGENAME=try-s3
PROJECTNAME=try-s3

## Docker target ------------------------

all: build run

.PHONY: build
build: ## build images of go1.11.0 + aws config
	docker build -t $(IMAGENAME) .

.PHONY: run
run: ## run go1.11.0 + aws config container
	docker run --name $(IMAGENAME) --rm -it -p 8081:8081 -v $(shell pwd):/go/src/${PROJECTNAME} $(IMAGENAME) bash

.PHONY: exec
exec: ## exec go1.11.0 + aws config container
	docker exec -it $(IMAGENAME) /bin/bash

.PHONY: help
help: ## show help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'