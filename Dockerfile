## Build
FROM golang:1.17.6-alpine3.14 AS build

WORKDIR /app

ENV DOCKER_BUILDKIT=0
ENV COMPOSE_DOCKER_CLI_BUILD=0

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
ENV CGO_ENABLED=0
RUN go build -o /main


## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /main /main

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/main"]
