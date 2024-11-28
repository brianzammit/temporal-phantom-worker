# Stage 1: Build the binary
FROM golang:1.23-alpine as builder

WORKDIR /app

# Copy the entire project, including the Makefile
COPY . .

# Build arguments
ARG VERSION
ARG COMMIT_HASH
ARG BUILD_TIME

# Run the build using the Makefile
RUN apk add --no-cache make && make build-linux

# Stage 2: Create the final image
FROM alpine:3.20

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/bin/linux/temporal-phantom .

# Set up the entrypoint
ENTRYPOINT ["/app/temporal-phantom"]

# Default command
CMD ["--help"]
