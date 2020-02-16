ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

COMMIT:=$(shell git rev-list -1 HEAD)
DATE:=$(shell date -uR)
GOVERSION:=$(shell go version | awk '{print $$3 " " $$4}')
GOOS:=$(shell go env GOOS)
GOARCH:=$(shell go env GOARCH)

DOCKER_IMAGE?=josefuentes/mr-wolf
DOCKER_IMAGE_TAG?=$(DOCKER_IMAGE):$(COMMIT)

define LDFLAGS
-X "github.com/j-fuentes/mr-wolf/pkg/version.Commit=$(COMMIT)" \
-X "github.com/j-fuentes/mr-wolf/pkg/version.Platform=$(GOOS)/$(GOARCH)" \
-X "github.com/j-fuentes/mr-wolf/pkg/version.BuildDate=$(DATE)" \
-X "github.com/j-fuentes/mr-wolf/pkg/version.GoVersion=$(GOVERSION)"
endef

GO_BUILD:=go build -ldflags '$(LDFLAGS)'
GO_INSTALL:=go install -ldflags '$(LDFLAGS)'

export GO111MODULE=on

.PHONY: build

build:
	cd $(ROOT_DIR) && $(GO_BUILD) -o builds/mr-wolf .

install:
	cd $(ROOT_DIR) && $(GO_INSTALL)

test:
	cd $(ROOT_DIR) && go test ./...

vet:
	cd $(ROOT_DIR) && go vet ./...

lint: vet
	cd $(ROOT_DIR) && golint

clean:
	cd $(ROOT_DIR) && rm -rf ./builds

# Docker image

build-docker-image:
	docker build --tag "$(DOCKER_IMAGE_TAG)" .

push-docker-image:
	docker tag $(DOCKER_IMAGE_TAG) $(DOCKER_IMAGE):latest
	docker push $(DOCKER_IMAGE_TAG)
	docker push $(DOCKER_IMAGE):latest
