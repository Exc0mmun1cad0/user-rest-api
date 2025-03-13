FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o app ./cmd/api/main.go


FROM alpine:3.21.3 AS runner

RUN apk update && apk upgrade

RUN adduser -D appuser
USER appuser

WORKDIR /app

COPY --from=builder /app/app .
COPY config config

EXPOSE 8080

CMD ["./app"]
