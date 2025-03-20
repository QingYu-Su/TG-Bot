# Build stage
FROM cgr.dev/chainguard/go:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o ./bin/bot .
RUN chmod +x ./bin/bot

# Run stage
FROM cgr.dev/chainguard/static:latest-glibc
WORKDIR /app
COPY ./config.yaml ./config.yaml
COPY --from=builder /app/bin/bot .
CMD ["./bot"]
