FROM golang:1.24.3-alpine AS builder

COPY . /github.com/Igorezka/auth/source/
WORKDIR /github.com/Igorezka/auth/source/

RUN go mod download
RUN go build -o ./bin/app/server cmd/grpc_server/main.go

FROM alpine:3.21.3

WORKDIR /root/
COPY --from=builder /github.com/Igorezka/auth/source/bin/app/server .

CMD ["./server"]