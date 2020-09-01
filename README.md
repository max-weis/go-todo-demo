# go-todo-demo

This is a demo project to demonstrate golang's `html/template` package.

required software:
* go
* docker/docker-compose
* make
* kubernetes
* helm

_The project only implements some parts of DDD. This project is just a demo and not production ready_

The project layout was inspired by a talk given at [Golang UK Conference 2016 by Marcus Olsson](https://www.youtube.com/watch?v=twcDf_Y2gXY). [Here](https://github.com/marcusolsson/goddd) you can find the corresponding GitHub Repository.

## run locally

Run `make` or `docker-compose up --build` to start the container

_the postgres container needs to be running before the demo application_, should the `go-todo` service shutdown, just restart the **compose file** or start `postgres` manually

## deploy to kubernetes via Helm

Further documentation can be found [here](./chart/README.md)

## run tests

Run `make ut` or `go test -v ./... -tags=unit` to run the unit tests