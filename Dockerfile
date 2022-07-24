FROM golang:1.18-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o server
RUN chmod +x scripts/wait-for