FROM golang:1.18-bullseye AS builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /app/gcp-stale-resource-finder

# final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/gcp-stale-resource-finder /app/gcp-stale-resource-finder
CMD [ "./gcp-stale-resource-finder" ]