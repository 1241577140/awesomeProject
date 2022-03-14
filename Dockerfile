FROM golang:1.16-alpine as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.io,direct \
    GIN_MODE=release \
    PORT=80

WORKDIR /

COPY . .
## 修改为任意目录下名称
RUN go build -o main /websocket/main.go

FROM alpine

WORKDIR /

COPY --from=builder /main .

EXPOSE 80

ENTRYPOINT ["./main"]