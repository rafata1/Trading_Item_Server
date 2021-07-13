FROM golang:1.16-buster as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN  go build -o main cmd/server/main.go
## Today ubuntu is using minimalized image by default, using ubuntu for better compatible than alpine
FROM ubuntu:20.04
WORKDIR /app/bin/
COPY --from=builder /app/main /app/bin/main
ENTRYPOINT ./main