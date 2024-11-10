# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -ldflags "-w" -o okx-oi-exporter .

# Stage 2: Create the final image
FROM gcr.io/distroless/static-debian12

WORKDIR /app

COPY --from=builder /app/okx-oi-exporter /app/

EXPOSE 8080

CMD ["/app/okx-oi-exporter"]