# Etap 1: Budowanie aplikacji
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Budowanie statyczne
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o swift-service ./cmd

# Etap 2: Obraz produkcyjny
FROM alpine:latest

WORKDIR /root/
RUN apk add --no-cache libc6-compat

COPY --from=builder /app/swift-service .

EXPOSE 8080
CMD ["./swift-service"]
