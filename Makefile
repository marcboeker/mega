.PHONY: ent graphql wire test

ent:
	go generate ./ent

graphql:
	go run github.com/99designs/gqlgen --config graph/gqlgen.yml

wire:
	wire ./app ./test

all: ent graphql wire

run:
	go run cmd/main.go

test:
	go test -v ./...