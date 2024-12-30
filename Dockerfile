# Use a lightweight version of Golang for building the application
FROM golang:alpine AS builder

# Set the working directory in the container
WORKDIR /src

# Set some environment variables
ENV CGO_ENABLED=0
ENV GOMODCACHE=/go/pkg/mod
ENV GOCACHE=/root/.cache/go-build

# Build the go app
RUN --mount=type=cache,target=$GOMODCACHE \
--mount=type=cache,target=$GOCACHE \
--mount=type=bind,target=. \
cd ./cmd/main && go build -ldflags="-w -s" -o /ttt

# Start a new stage from scratch for the final image
FROM alpine:latest AS final

# Set the working directory in the container
WORKDIR /

# Set the timezone and install CA certificates
RUN --mount=type=cache,target=/var/cache/apk \
apk --update add ca-certificates tzdata && update-ca-certificates

# Expose port
EXPOSE 8080

# Copy the pre-built binary file from the previous stage
COPY --from=builder /ttt .

# Command to run the executable
ENTRYPOINT ["./ttt"]
