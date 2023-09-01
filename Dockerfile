FROM golang:latest

RUN go version
ENV GOPATH=/

RUN apt-get update
RUN apt-get -y install postgresql-client
COPY ./ ./

RUN go mod download
RUN go build -o avito-test ./cmd/segmenter/main.go

EXPOSE 8080 8080