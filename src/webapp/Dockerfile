# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.9 as builder

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/phillikus/goTherm/src/webapp

WORKDIR /go/src/github.com/phillikus/goTherm/src/webapp

RUN go get github.com/lib/pq && go build -o main

# Document that the service listens on port 5000.
EXPOSE 5000

CMD ["./main"]