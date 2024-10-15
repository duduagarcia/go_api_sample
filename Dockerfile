# Stage 1: Build the Go application
FROM golang:1.21-alpine AS builder
WORKDIR /app

# Install Git and build dependencies (such as GCC)
RUN apk add --no-cache git gcc musl-dev

# Copy go mod and sum files
COPY go.mod go.sum ./

RUN go mod download

# Copy the source code
COPY . .


# Set environment variables for CGO and target architecture
ENV CGO_ENABLED=1 GOOS=linux

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init

# Build the Go app for Linux with CGO enabled
RUN go build -o main .


# Stage 2: Copy the binary and run it in a smaller image
FROM alpine:latest
WORKDIR /app

# Install build dependencies for CGO (SQLite support)
RUN apk add --no-cache postgresql-client
# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
