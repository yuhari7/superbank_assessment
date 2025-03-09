# Use official Golang image as builder
FROM golang:1.21 AS builder

WORKDIR /app

# Copy go.mod and go.sum files first to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN go build -o main cmd/main.go

# Use a lightweight image for final deployment
FROM alpine:latest

WORKDIR /root/

# Copy built binary from builder stage
COPY --from=builder /app/main .

# Expose port 3000
EXPOSE 3000

# Run the binary
CMD ["./main"]
