# Base image for building the Go application
FROM golang:1.21-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the source code
COPY . .

# Build the Go application
RUN go mod tidy
RUN go build -o /quote-service main.go

# --- Final image (smaller for production) ---
FROM alpine:latest
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /quote-service .

# Expose the port
EXPOSE 8080

# Command to run the executable
CMD ["./quote-service"]
