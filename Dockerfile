FROM golang:1.14-alpine as build

WORKDIR $GOPATH/app/

RUN apk add git

# copy and download dependencies
#COPY go.* ./
#RUN go mod download

#compile app
COPY . .

RUN mkdir -p log

RUN go run cmd/gotodo/main.go
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main
#RUN chmod +x main
#ENTRYPOINT [ "./main" ]
#resulting app
#FROM alpine:3.12 as final
#FROM scratch as final
#COPY --from=build go/app/main /app/
#WORKDIR /app
#RUN mkdir log && chmod +x main

#ENTRYPOINT [ "./main" ]