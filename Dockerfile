FROM golang:alpine AS builder

WORKDIR /tmp/build

COPY . .

RUN sed -i 's/http:\/\/dl-cdn.alpinelinux.org/https:\/\/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk --no-cache add tzdata \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod download \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o kns

FROM scratch

COPY --from=builder /tmp/build/kns /kns
COPY --from=builder /usr/share/zoneinfo/PRC /etc/localtime

EXPOSE 60080

ENTRYPOINT ["/kns"]
