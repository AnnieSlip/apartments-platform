
FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download

# copy the rest of the project
COPY . .

# Build the precompute binary
RUN go build -o precompute cmd/precompute/main.go

# run stage
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/precompute .
CMD ["./precompute"]
