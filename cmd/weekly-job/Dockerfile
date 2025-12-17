# ------------------------
# Build stage
# ------------------------
FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o weekly-job cmd/weekly-job/main.go

# ------------------------
# Run stage
# ------------------------
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/weekly-job .

CMD ["./weekly-job"]
