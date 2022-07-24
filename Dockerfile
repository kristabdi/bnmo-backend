FROM golang:1.18-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o server

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait .
RUN chmod +x wait

CMD /app/wait && /app/server