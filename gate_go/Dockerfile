# Builer образ для сборки.
# Требуется искулючитально для компиляции исполняемого файла
FROM golang:1.19 AS builder
WORKDIR /service
COPY . .
RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags '-extldflags "-static"' -o /gateway ./cmd/gateway/main.go

# Service образ без артифактов и сборочного мусора..
FROM alpine:latest
COPY --from=builder /gateway /bin/
ENTRYPOINT ["gateway"]