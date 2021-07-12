FROM golang:1.15 AS builder
WORKDIR /go/src/ws-server/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ws-server .

FROM alpine:latest
WORKDIR /root/ws-server
COPY --from=builder /go/src/ws-server .
CMD ["./ws-server"]
