# Etap 1: Budowanie aplikacji
FROM golang:1.21 AS builder

# Ustaw katalog roboczy
WORKDIR /app

# Skopiuj pliki modułów i pobierz zależności
COPY go.mod go.sum ./
RUN go mod download

# Skopiuj resztę projektu
COPY . .

# Buduj aplikację – zakładamy, że entrypoint jest w cmd/main.go
RUN go build -o swift-service ./cmd

# Etap 2: Obraz produkcyjny
FROM alpine:latest

WORKDIR /root/

# Skopiuj binarkę z etapu build
COPY --from=builder /app/swift-service .

# Otwórz port – zmień jeśli aplikacja nasłuchuje na innym
EXPOSE 8080

# Uruchom aplikację
CMD ["./swift-service"]