FROM golang:1.19 as builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o ordersystem ./cmd/ordersystem

FROM scratch
COPY --from=builder /app/ordersystem .
COPY --from=builder /app/cmd/ordersystem/.env .
CMD ["./ordersystem"]
