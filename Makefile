default: run

VERSION ?= v1.0
APP ?= go-todo-demo
DOCKER_IMAGE ?= maxweis/${APP}

ut:
	@go test -v ./... -tags=unit

#it:
#	@go test ./... -tags=integration

run:
	@docker-compose up --build

docker-build:
	@docker image build -t $(DOCKER_IMAGE):$(VERSION) .

docker-push:
	@docker image push $(DOCKER_IMAGE):latest
	@docker image push $(DOCKER_IMAGE):$(VERSION)