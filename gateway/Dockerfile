FROM golang:1.23.5-bullseye AS builder

WORKDIR /app

COPY . .

RUN go work sync

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o ./gateway/bin/gateway ./gateway/main.go

FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=builder /app/gateway/bin/gateway /app/gateway

CMD ["/app/gateway"]