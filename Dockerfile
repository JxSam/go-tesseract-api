ARG GO_VERSION="1.26"

# =========================
# BUILD
# =========================
FROM golang:${GO_VERSION}-bookworm AS builder

WORKDIR /build

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        libtesseract-dev \
        libleptonica-dev \
        pkg-config \
        gcc \
        g++ \
    && rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux \
    go build -ldflags="-s -w" \
    -o /tesseract-api \
    ./cmd/main.go


# =========================
# RUNTIME
# =========================
FROM debian:bookworm-slim

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        tesseract-ocr \
        tesseract-ocr-rus \
        tesseract-ocr-eng \
        poppler-utils \
        ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /tesseract-api /app/tesseract-api

RUN mkdir -p /app/config
COPY --from=builder /build/config/local.yaml /app/config/local.yaml

ENV CONFIG_PATH=/app/config/local.yaml

EXPOSE 4560

CMD ["./tesseract-api"]
