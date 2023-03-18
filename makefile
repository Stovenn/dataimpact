fmt:
	go fmt ./...

lint:fmt
	golint ./...

vet:
	go vet ./...

build:vet
	go build -o bin/dataimpact cmd/main.go

run:
	./bin/dataimpact

test:
	go test -v -cover ./...

mongo:
	docker run -d --name mongo -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=secret -dp 27017:27017 mongo

.PHONY: fmt lint vet build run test mongo