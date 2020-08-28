FROM golang:1.14-alpine as build

WORKDIR $GOPATH/app/

RUN apk add git

# copy and download dependencies
COPY go.* .
RUN go mod download

#compile app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main && \
    mkdir -p log

#resulting app
FROM scratch as final
COPY --from=build go/app/main /app/
COPY --from=build go/app/static /app/static
COPY --from=build go/app/log/ /app/log
WORKDIR /app
EXPOSE 8080
ENTRYPOINT [ "./main" ]