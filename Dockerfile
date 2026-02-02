FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o candipack cmd/server/main.go

FROM alpine:3.19

# minimal chromium setup
RUN apk add --no-cache chromium nss font-noto

WORKDIR /app
COPY --from=builder /app/candipack .
COPY --from=builder /app/templates ./templates

ENV CHROME_PATH=/usr/bin/chromium-browser

CMD ["./candipack", "server"]
