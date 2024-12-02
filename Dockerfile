FROM golang:1.23-alpine AS builder

WORKDIR /app/
COPY go.* ./
RUN go mod download

COPY ./ ./

RUN go build -o main ./app/main.go

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /app/main ./

COPY .env .

EXPOSE 8080

ENTRYPOINT ["./main"]