# Step 1: Build Stage
FROM golang:1.23.1-alpine as builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory inside the container
WORKDIR /app

# Install git needed for go modules
RUN apk add --no-cache git

# COPY go.mod go.sum first for dependency caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Install swag for Swagger generation
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy whole project
COPY . . 

# Generate Swagger docs (this will create ./docs)
RUN swag init -g main.go -o ./docs

# Build the Go app
RUN go build -o main

# Step 2: Run stage
FROM alpine:latest

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/main .

# Expose the port 
EXPOSE 8080

# RUN the application
CMD ["./main"]