FROM golang:1.21-alpine AS builder

ARG APP_NAME

WORKDIR /app

RUN go mod init app
RUN go get github.com/pressly/goose/v3
RUN go get github.com/lib/pq
RUN go get github.com/jmoiron/sqlx

COPY ./sql ./sql
COPY ./cmd ./cmd

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/bin ./cmd

# Path: Dockerfile
FROM alpine:3.18.4

RUN apk update && apk --no-cache add tzdata

COPY --from=builder /app/bin /app/bin
COPY --from=builder /app/sql /app/sql

ENTRYPOINT ["/app/bin"]
