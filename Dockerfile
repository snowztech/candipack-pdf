FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o candipack-pdf ./cmd/server

FROM alpine:3.19
RUN apk add --no-cache chromium font-noto
WORKDIR /app
COPY --from=builder /app/candipack-pdf .
COPY --from=builder /app/templates ./templates
ENV CHROME_PATH=/usr/bin/chromium-browser
ENV PORT=9000
EXPOSE 9000
CMD ["./candipack-pdf"]
