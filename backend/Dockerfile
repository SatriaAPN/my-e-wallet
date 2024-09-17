# Build stage
FROM golang:1.22-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go modules manifest files
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go app
RUN go build -o backend .

# Final stage: Run the binary
FROM alpine:3.18
WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/backend /app/backend

# Expose the backend port
EXPOSE 8080

# Run the binary
CMD ["/app/backend"]
