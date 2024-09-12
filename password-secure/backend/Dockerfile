# backend/Dockerfile
FROM golang:1.23-alpine

WORKDIR /app

# Copia los archivos al contenedor
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compila la aplicación
RUN go build -o main ./cmd/main.go

EXPOSE 8080

# Ejecuta la aplicación
CMD ["./main"]
