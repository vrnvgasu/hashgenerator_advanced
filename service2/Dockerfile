#---Build stage---
FROM golang:1.22 AS builder
COPY . /go/src/
WORKDIR /go/src/cmd/service2

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags='-w -s' -o /go/bin/service

#---Final stage---
FROM alpine:latest
COPY --from=builder /go/bin/service /go/bin/service
COPY --from=builder /go/src/migrations /go/bin/migrations
CMD /go/bin/service --host '0.0.0.0'