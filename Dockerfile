# Build stage
# Using Go 1.23 to support dependencies requiring Go 1.23+
FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

# Copy source code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s -extldflags '-static'" \
    -trimpath \
    -o /app/vkr \
    ./cmd/vkr

FROM gcr.io/distroless/static-debian11:nonroot

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /app/vkr /app/vkr

COPY --from=builder /build/config /app/config

WORKDIR /app

USER nonroot:nonroot

EXPOSE 8081

ENTRYPOINT ["/app/vkr"]

