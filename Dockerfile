# Use the official Golang image as the base
FROM golang:1.21-alpine

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create working directory
WORKDIR /app

# Copy go.mod and go.sum files first
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go app
RUN go build -o server main.go

# Expose port
EXPOSE 3000

# Run the server
CMD ["./server"]
