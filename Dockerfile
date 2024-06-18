FROM golang:1.22-alpine AS builder

RUN apk update && \
    apk add ca-certificates && \
    apk add --no-cache git

WORKDIR /go/src/app

COPY . .

RUN GOOS=linux GOARCH=arm64 go build -a -o /usr/local/bin/chat-copilot ./cmd/...

FROM alpine:latest

WORKDIR /app

COPY --from=builder /usr/local/bin/chat-copilot /usr/local/bin/chat-copilot

RUN chmod +x /usr/local/bin/chat-copilot 

ENTRYPOINT [ "chat-copilot" ]