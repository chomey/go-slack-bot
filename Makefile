VERSION := $(shell cat VERSION)+$(shell git log -1 --pretty=format:%h)
IMAGE_TAG = $(shell echo $(VERSION) | sed 's|[+:]|-|g')
SERVICE_NAME := go_slack_bot
IMAGE_NAME := chomey/$(SERVICE_NAME):$(IMAGE_TAG)
RC_HOSTNAME ?= localhost

build:
#	bash -c "CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -ldflags '-X github.com/chomey/go-slack-bot/service.VERSION=$(VERSION)' -o bin/go_slack_bot_darwin_amd64"
#	bash -c "CGO_ENABLED=0 GOARCH=386 GOOS=linux go build -ldflags '-X github.com/chomey/go-slack-bot/service.VERSION=$(VERSION)' -o bin/go_slack_bot_linux_386"
	bash -c "CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags '-X github.com/chomey/go-slack-bot/service.VERSION=$(VERSION)' -o bin/go_slack_bot_linux_amd64"
#	bash -c "CGO_ENABLED=0 GOARCH=386 GOOS=windows go build -ldflags '-X github.com/chomey/go-slack-bot/service.VERSION=$(VERSION)' -o bin/go_slack_bot_windows_386"
#	bash -c "CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -ldflags '-X github.com/chomey/go-slack-bot/service.VERSION=$(VERSION)' -o bin/go_slack_bot_windows_amd64"
	docker build -t "chomey/go-slack-bot:dev" .

run:
	docker-compose rm -f 2>/dev/null || true
	VERSION=$(IMAGE_TAG) RC_HOSTNAME=$(RC_HOSTNAME) RC_PRIVATE_KEY=$(RC_PRIVATE_KEY) docker-compose up

local_test:
	go test `go list ./... | grep -v /vendor/`

image: build
	docker build -t $(IMAGE_NAME) . && docker tag $(IMAGE_NAME) "go-slack-bot:localdev"
clean:
	docker rm -f $(BUILD_CONTAINER_NAME) 2> /dev/null || true
	rm bin/* || true

.PHONY: build run local_run build clean container_

