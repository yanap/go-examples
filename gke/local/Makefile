# Makefile

GO=go 

local:
	docker-compose stop
	docker-compose rm -f
	docker-compose build
	docker-compose up -d
	docker-compose ps

clean:
	docker container prune && docker images prune
	docker rmi $(docker images -q)

build:
	$(GO) build -o app main.go

test:
	$(GO) test ./... -v

fmt:
	$(GO) fmt ./...
