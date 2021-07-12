FROM golang:1.15 AS builder
WORKDIR /go/src/resful-server/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o resful-server .

FROM alpine:latest
WORKDIR /root/resful-server
COPY --from=builder /go/src/resful-server .
CMD ["./resful-server"]
