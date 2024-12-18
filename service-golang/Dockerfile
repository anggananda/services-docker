# Gunakan base image untuk Golang
FROM golang:1.20-alpine

# Set working directory di dalam container
WORKDIR /app

COPY .env .
# Copy semua file ke dalam container
COPY . .

# Unduh dependencies
RUN go mod tidy

# Build aplikasi
RUN go build -o service-golang

# Jalankan aplikasi
CMD ["./service-golang"]
