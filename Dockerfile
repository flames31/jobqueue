# Build stage
FROM golang:1.23.3-alpine AS builder

WORKDIR /app

# Install git and other needed packages
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o jobqueue ./cmd/server

# Run stage
FROM alpine:latest

WORKDIR /app

# Add SSL certs for HTTPS calls, if needed
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/jobqueue .

# Copy static files or .env if needed
COPY ./cmd/server/.env .

EXPOSE 8080

CMD ["./jobqueue"]
