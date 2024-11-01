run:
	go run cmd/main.go

get-deps:
	go mod tidy

generate-swagger:
	swag init -g cmd/main.go