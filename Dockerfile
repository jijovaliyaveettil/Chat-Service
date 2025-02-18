# Use official Golang image as base
FROM golang:1.21

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create working directory
WORKDIR /app

# Copy Go modules manifests
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Build the application
RUN go build -o chat-service cmd/server/main.go

# Expose the port for WebSockets
EXPOSE 8080

# Run the application
CMD ["./chat-service"]