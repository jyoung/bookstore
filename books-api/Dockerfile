# --- Build Stage ---
FROM golang:1.26-alpine AS build
WORKDIR /app

# Copy dependency manifests and download layers
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/books-api .

# --- Run Stage ---
FROM alpine:latest
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/bin/books-api .

# Expose the Go Chi server port
EXPOSE 8080

CMD ["./books-api"]
