IMAGENAME=try-s3
PROJECTNAME=try-s3

## Docker target ------------------------

all: build run

.PHONY: build
build: ## build images of go-1.6.2 environment
	docker build -t $(IMAGENAME) .

.PHONY: run
run: ## run container of go-1.6.2 environment
	docker run --name $(IMAGENAME) --rm -it -p 8081:8081 -v $(shell pwd):/go/src/${PROJECTNAME} $(IMAGENAME) bash

.PHONY: exec
exec: ## exec go1.6.2 container
	docker exec -it $(IMAGENAME) /bin/bash

.PHONY: stop
stop: ## stop container.
	docker stop $(IMAGENAME)

.PHONY: rmi
rmi: ## remove docker image.
	docker rmi -f $(IMAGENAME)

.PHONY: clean
clean: stop rmi ## clean up docker image & container.