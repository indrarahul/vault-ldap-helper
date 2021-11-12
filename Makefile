BUILDPATH=$(CURDIR)
GO=$(shell which go)
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GORUN=$(GO) run

build:
	$(GOBUILD) -o bin/vault_ldap_helper
	
run:
	$(GORUN) main.go -config=.config.yaml -verbose=0

clean:
	$(GOCLEAN)
