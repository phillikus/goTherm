# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.9 as builder

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/phillikus/goTherm/src/webapp

WORKDIR /go/src/github.com/phillikus/goTherm/src/webapp

RUN go get github.com/lib/pq && CGO_ENABLED=0 GOOS=linux go build -o main

FROM alpine:latest as runner

# TODO: remove curl and use custom go program for health check in docker-compose.yml
RUN apk --no-cache add curl

COPY --from=builder /go/src/github.com/phillikus/goTherm/src/webapp .

# Document that the service listens on port 5000.
EXPOSE 5000


ENTRYPOINT ["./main", "-prod"]