# golang image to download dependencies and build the project
FROM golang:1.9 as builder

ADD . /go/src/github.com/phillikus/goTherm/src/webapp

WORKDIR /go/src/github.com/phillikus/goTherm/src/webapp

RUN go get github.com/lib/pq && CGO_ENABLED=0 GOOS=linux go build -o main

# alpine image to run the compiled app
FROM alpine:latest as runner

# install curl for healthcheck
RUN apk --no-cache add curl

COPY --from=builder /go/src/github.com/phillikus/goTherm/src/webapp .

# env variables for development
ENV CONNECTION_STRING host=postgres port=5432 user=db_user password=Start123 dbname=therm_db sslmode=disable

EXPOSE 5000

ENTRYPOINT ["./main", "-prod"]