# syntax=docker/dockerfile:1

FROM golang:1.24.2-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/goshorty ./cmd/main.go

FROM alpine:3.21

RUN adduser -D -H -u 65532 appuser
WORKDIR /app
COPY --from=build /out/goshorty /app/goshorty

ENV HTTP_ADDR=:8080 \
    METRICS_ADDR=127.0.0.1:6060 \
    PUBLIC_BASE_URL=http://localhost:8080

EXPOSE 8080

USER appuser
CMD ["/app/goshorty"]
