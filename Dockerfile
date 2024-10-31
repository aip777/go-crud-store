# Stage 1: Build GO
FROM golang:1.23-alpine AS builder
LABEL authors="palash"

# Install build dependencies
RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
# Copy the remaining source code
COPY . .
RUN go build -o main .

# Stage 2: Run the application
FROM alpine:latest
WORKDIR /root

# Copy from the builder stage
COPY --from=builder /app/main .
COPY .env .env
COPY ./data /data
VOLUME /data
EXPOSE 8000
ENV DATABASE_PATH=/data/store.db
CMD ["./main"]
