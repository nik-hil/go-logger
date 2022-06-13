ARG GO_VERSION=1.18

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
# RUN go build -o ./app  ./src/*
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./app  ./src/*

FROM alpine:latest

# RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api
COPY --from=builder /api/app .
# COPY --from=builder /api/test.db .

ENV GIN_MODE release

EXPOSE 8080

ENTRYPOINT ["./app"]