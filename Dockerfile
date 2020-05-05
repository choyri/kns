FROM golang:alpine AS builder

WORKDIR /tmp/build

COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod download \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o kns

FROM golang:alpine

COPY --from=builder /tmp/build/kns /kns

EXPOSE 60080

ENTRYPOINT ["/kns"]
