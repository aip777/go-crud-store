FROM golang:1.20-alpine AS builder
LABEL authors="palash"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .


FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
VOLUME /data
COPY ./data /data
COPY .env .env
EXPOSE 8000
ENV DATABASE_PATH=/data/store.db
CMD ["./main"]