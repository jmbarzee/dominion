
.PHONY: dominion domain build-dominion build-domain dominion-gen-config domain-gen-config test

DOMINION_CONFIG ?= "env/dominion.env"
DOMAIN_CONFIG ?= "env/domain.env"

# Runs
dominion: build-dominion
	( source $(DOMINION_CONFIG); ./bin/dominion )

domain: build-domain
	( source $(DOMAIN_CONFIG); ./bin/domain )


# Builds
build-dominion:
	go build -o bin/dominion cmd/dominion/main.go

build-domain:
	go build -o bin/domain cmd/domain/main.go


# Config Generation
dominion-gen-config:
	go build -o bin/dominion-gen-config cmd/dominion/genconfig/main.go
	./bin/dominion-gen-config $(DOMINION_CONFIG)

domain-gen-config:
	go build -o bin/domain-gen-config cmd/domain/genconfig/main.go
	./bin/domain-gen-config $(DOMAIN_CONFIG)


# Unit Tests
test: 
	go test ./...

