# Build Stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/main.go

# Run Stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /main .

EXPOSE 3000

ENTRYPOINT ["/app/main"]