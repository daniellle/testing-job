# Build stage
FROM golang:1.20 as builder
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code and build the binary
COPY . .
RUN go build -o app .

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .

# Command to run the Go binary
CMD ["./app"]
