build:
	docker build -t plugin-builder ./

run:
	docker run --rm -v ${PWD}:/home/gradle/src/ \
	plugin-builder  /bin/sh -c "${CMD}"
