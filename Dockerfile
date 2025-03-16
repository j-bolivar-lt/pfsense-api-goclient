FROM golang:1.21.7-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod tidy

COPY . .

RUN go build -o main ./pfsenseapi

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]
