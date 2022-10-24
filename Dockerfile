#build stage
FROM golang:alpine AS builder
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk -u --no-cache add git
WORKDIR /go/src/app
COPY . .
RUN go build -o /go/bin/app -v main.go

#final stage
FROM alpine:latest
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /chat/app
WORKDIR /chat
ENTRYPOINT /chat/app