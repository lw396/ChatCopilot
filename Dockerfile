FROM golang:1.23-alpine AS builder

RUN apk add --no-cache ca-certificates git gcc make musl-dev

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN cd pkg/audio/silk && make

RUN CGO_ENABLED=1 go build -o chat-copilot ./cmd/*.go

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /go/src/app/chat-copilot /usr/local/bin/chat-copilot

ENTRYPOINT [ "chat-copilot" ]