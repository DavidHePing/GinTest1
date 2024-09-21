# Stage 1: Build the Go application
FROM golang:1.21.5 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# Stage 2: Create a smaller image for running the application
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go binary from the builder stage
COPY --from=builder /app/myapp .
# Copy the config.json file
COPY --from=builder /app/config.json .

# Expose the port on which the application runs (modify as needed)
EXPOSE 8080

# Command to run the application
ENTRYPOINT  ["./myapp"]
