# Use a lightweight version of Golang for building the application
FROM golang:alpine AS builder

# Set the working directory in the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Download all dependencies
RUN go mod download

# Build the Go app
RUN cd ./cmd/main && CGO_ENABLED=0 GOOS=linux go build -o /main

# Start a new stage from scratch for the final image
FROM alpine:latest

WORKDIR /

# Copy the pre-built binary file from the previous stage
COPY --from=builder /main .

# Set the timezone and install CA certificates
RUN apk --no-cache add ca-certificates tzdata

# Expose port
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["./main"]
