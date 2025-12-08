.PHONY: run build test lint docker-build docker-run clean

run:
	go run main.go

build:
	go build -o bin/kanban-backend main.go

test:
	go test -v -cover ./...

lint:
	golangci-lint run

docker-build:
	docker build -t kanban-backend:latest .

docker-run:
	docker run -p 8080:8080 kanban-backend:latest

clean:
	rm -rf bin/
	go clean
