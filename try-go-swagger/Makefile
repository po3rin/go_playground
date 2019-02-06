PHONY: validate

validate:
	swagger validate ./swagger.yml
run:
	go run cmd/todo-list-server/main.go
generate:
	swagger generate server -A TodoList -f ./swagger.yml
