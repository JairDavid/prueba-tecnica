FROM golang:1.22 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
    
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
WORKDIR /app/cmd
RUN go build -o taskMicroservice

WORKDIR /app

EXPOSE 9000
ENTRYPOINT [ "./cmd/taskMicroservice" ]