# Base Build
FROM golang:1.26-alpine AS base
WORKDIR /app
COPY src/go.mod src/go.sum ./
RUN go mod download
COPY ./src .

# Dev Build
FROM base as dev
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN CGO_ENABLED=0 go build -gcflags="all=-N -l" -o /app/books-api .
EXPOSE 8080 40000
CMD ["dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "./books-api", "--continue"]

# Production Build
FROM base as builder
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/books-api .

# Final Production Build
FROM alpine:latest as prod
WORKDIR /app
COPY --from=builder /app/books-api .
EXPOSE 8080
CMD ["./books-api"]
