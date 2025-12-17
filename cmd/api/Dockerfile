FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the API binary
RUN go build -o server ./cmd/api/main.go

EXPOSE 8080

# Run the server
CMD ["./server"]
