# Build stage
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest do projeto
COPY . .

# Build o app (compilando o diretório, não o arquivo)
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/app

# Final stage
FROM alpine:latest

WORKDIR /app

# Copia o binário e migrations
COPY --from=builder /app/main .
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080

CMD ["./main"]