# syntax=docker/dockerfile:1

# Stage 1: Compile the updated Go application with Places Autocomplete & CRM Sync
FROM golang:1.23-bookworm AS builder
WORKDIR /app
COPY gmaps-src/go.mod gmaps-src/go.sum ./
RUN go mod download
COPY gmaps-src/ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/google-maps-scraper .

# Stage 2: Use the pre-configured official image with Chromium and Playwright
FROM gosom/google-maps-scraper:v1.18.1

# Replace default binary with our updated binary
COPY --from=builder /app/google-maps-scraper /usr/bin/google-maps-scraper

ENV PORT=8080
EXPOSE 8080

ENTRYPOINT ["sh", "-c", "google-maps-scraper -web -addr 0.0.0.0:${PORT:-8080} -data-folder /gmapsdata"]
