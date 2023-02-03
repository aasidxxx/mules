#FROM golang:1.14-buster

FROM golang:1.19

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o mu-app ./cmd/server/main.go

CMD ["./mu-app"]