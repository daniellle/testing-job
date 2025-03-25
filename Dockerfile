# Build stage
FROM golang:1.20 as builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app .

# Final image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
RUN apk add --no-cache postgresql-client

CMD ["./app"]
