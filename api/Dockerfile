FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./pkg ./pkg
COPY ./internal ./internal

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/bin ./internal/cmd/main

# Path: Dockerfile
FROM alpine:3.18.4 AS runner

RUN apk update && apk --no-cache add tzdata

COPY --from=builder /app/bin /app/bin

ENTRYPOINT ["/app/bin"]
