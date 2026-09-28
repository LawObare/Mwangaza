# syntax=docker/dockerfile:1
# Mwangaza backend — multi-stage build: compile a static Go binary, then run
# it on a minimal Alpine image as an unprivileged user.

FROM golang:1.27-alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0

# Cache module downloads separately from source changes.
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags="-s -w" -o /out/mwangaza-api ./cmd

FROM alpine:3.22
# ca-certificates is required for the outbound HTTPS call to OpenWeatherMap.
RUN apk add --no-cache ca-certificates \
 && adduser -D -H -u 10001 mwangaza \
 && mkdir -p /data \
 && chown mwangaza:mwangaza /data

COPY --from=build /out/mwangaza-api /usr/local/bin/mwangaza-api

USER mwangaza

# Defaults; backend/.env is supplied at runtime via compose `env_file`.
# Compose overrides DATABASE_PATH to point inside the /data volume.
ENV PORT=8080 \
    DATABASE_PATH=/data/lakenet.db \
    USE_MOCK_DATA=true

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget -q -O /dev/null http://127.0.0.1:8080/api/health || exit 1

ENTRYPOINT ["/usr/local/bin/mwangaza-api"]