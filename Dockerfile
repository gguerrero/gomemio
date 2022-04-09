# syntax=docker/dockerfile:1
FROM golang:1.17-alpine

WORKDIR /app
RUN cd /app

COPY go.mod ./
# COPY go.sum ./

RUN go mod download

COPY *.go .
COPY cmd cmd
COPY commands commands
COPY data data
COPY scanner scanner
COPY service service

RUN go build -o gomemio cmd/main.go
RUN chmod +x gomemio

EXPOSE 8080

ENTRYPOINT [ "/app/gomemio" ]
CMD [ "-p", "8080" ]
