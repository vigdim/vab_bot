
.PHONY: run build clean test help default

GOFLAGS ?= $(GOFLAGS:)

BIN_NAME=vab_bot
GOPRIVATE_NAME="github.com/vigdim/vab_bot"

PROJECT_NAME := "vab_bot"
PKG := "https://github.com/vigdim/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)


VERSION = $(shell grep 'const version' version.go | sed -E 's/.*"(.+)"$$/v\1/')
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')
IMAGE_NAME := kktlecensebot/${BIN_NAME}

PREV_IMAGES=$(shell docker images --filter=reference='$(IMAGE_NAME)*' -q | uniq | tr '\n' ' ')
PREV_CONTAINERS=$(shell docker ps -qa --no-trunc --filter 'name=$(BIN_NAME)*' -q | uniq | tr '\n' ' ')

default: help

help-alt: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


help:
	@echo 'Management commands:'
	@echo
	@echo 'Usage:'
	@echo '    make run             Run the project with tests.'
	@echo '    make build           Compile the project.'
	@echo '    make bin-run         Compile and run binary the project.'
	@echo '    make bench           Bench the project.'
	@echo '    make docker-build    Compile optimized for alpine linux.'
	@echo '    make docker-up       Running docker image for alpine linux.'
	@echo '    make tag             Tag image created by package with latest, git commit and version'
	@echo '    make test            Run tests on a compiled project.'
	@echo '    make push            Push tagged images to registry'
	@echo '    make clean           Clean the directory tree.'
	@echo

coverage: ## Generate global code coverage report
	./tools/coverage.sh;

coverhtml: ## Generate global code coverage report in HTML
	./tools/coverage.sh html;


run: ## Run the project with tests.
	go run $(GOFLAGS) ./main.go

install:
	@echo "\e[34;1mInstall vendors...\e[0m"
	@go mod download
	@echo "\e[36;1mDone...\e[0m \n"

build: ## Compile the project.
	@echo "\e[32;1mBuilding ${BIN_NAME} ${VERSION}\e[0m"
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/${BIN_NAME}
	chmod +x ./bin/${BIN_NAME}
	@echo "\e[36;1mDone...\e[0m \n"

bin-run: build ## Compile and run binary the project.
	@echo "\e[32;1mExecuting ${BIN_NAME} ${VERSION}\e[0m"
	./bin/${BIN_NAME}

docker-rmi:
ifneq ($(PREV_IMAGES),)
	@echo "\e[33;1mDeleting previous buld images ${IMAGE_NAME}\e[0m"
	docker rmi -f $(PREV_IMAGES)
	@echo "\e[36;1mDone...\e[0m \n"
endif

docker-build: docker-rmi
	@echo "\e[33;1mBuilding image ${IMAGE_NAME}\e[0m"
	docker build -t ${IMAGE_NAME}:$(GIT_COMMIT) .
	@echo "\e[36;1mDone...\e[0m \n"

docker-stop:
ifneq ($(PREV_CONTAINERS),)
	@echo "\e[32;1mStoping image ${IMAGE_NAME}\e[0m"
	docker container stop $(PREV_CONTAINERS)
	docker rm $(PREV_CONTAINERS)
	@echo "\e[36;1mDone...\e[0m \n"
endif


docker-up:
	@echo "\e[32;1mStarting container from image ${IMAGE_NAME}-$(GIT_COMMIT)\e[0m"
	docker run -d -p 8091:3000 -e APP_ENV=release --restart=always \
#	-v ${PWD}/config:/root/config -v ${PWD}/data:/root/data \
	--name ${BIN_NAME} ${IMAGE_NAME}:$(GIT_COMMIT)
	@echo "\e[36;1mDone...\e[0m \n"

docker-log:
	docker container logs $(PREV_CONTAINERS) -f

test: install
	@echo -e "\e[32;1mStarting tests...\e[0m"
	go test $(GOFLAGS) ./...

bench: install ## Bench the project.
	@echo "\e[32;1mStarting benchmarks...\e[0m"
	go test -run=NONE -bench=. $(GOFLAGS) ./...

clean: ## Clean the directory tree.
	@echo "\e[35;1mCleaning previous build...\e[0m"
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}
	go clean $(GOFLAGS) -i ./...
	@echo "\e[36;1mDone...\e[0m \n"

tag: docker-build
	@echo "Tagging: latest ${VERSION} $(GIT_COMMIT)"
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):$(GIT_COMMIT)
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):${VERSION}
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):latest

push: tag
	@echo "Pushing docker image to registry: latest ${VERSION} $(GIT_COMMIT)"
	docker push $(IMAGE_NAME):$(GIT_COMMIT)
	docker push $(IMAGE_NAME):${VERSION}
	docker push $(IMAGE_NAME):latest
