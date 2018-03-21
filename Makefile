# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=countme
BINARY_UNIX=$(BINARY_NAME)_unix
DOCKERCMD=docker
DOCKERREPO=mfouilleul
DOCKERBUILD=$(DOCKERCMD) build
DOCKERPUSH=$(DOCKERCMD) push
DOCKERTAG=$(DOCKERCMD) tag

VERSION = $(shell cat ./VERSION)
LDFLAGS = -ldflags "-X main.version=${VERSION}"

all: deps build docker

build:
	$(GOBUILD) ${LDFLAGS} -o $(BINARY_NAME)
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) ${LDFLAGS} -o $(BINARY_NAME) ./...
	./$(BINARY_NAME)
deps:
	$(GOGET) github.com/go-redis/redis
	$(GOGET) net/http
	$(GOGET) github.com/ghodss/yaml
	$(GOGET) github.com/cenkalti/backoff
	$(GOGET) github.com/sirupsen/logrus



docker: docker-build docker-push

docker-build:
	$(DOCKERBUILD) -t $(DOCKERREPO)/$(BINARY_NAME):latest .
	$(DOCKERTAG) $(DOCKERREPO)/$(BINARY_NAME):latest $(DOCKERREPO)/$(BINARY_NAME):$(VERSION)

docker-push:
	$(DOCKERPUSH) $(DOCKERREPO)/$(BINARY_NAME):latest
	$(DOCKERPUSH) $(DOCKERREPO)/$(BINARY_NAME):$(VERSION)
