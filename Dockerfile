# First stage: Build Go binary
FROM golang:1.17 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Second stage: Create a minimal image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/app .

# Expose the port the app runs on
EXPOSE 8888

# Command to run the executable
CMD ["./app"]
