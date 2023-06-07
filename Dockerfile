FROM golang:1.20

WORKDIR /app
ADD . /app
RUN go env -w GOPROXY=https://goproxy.cn,direct && \
    go mod tidy && \
    go build -o server

EXPOSE 80

ENTRYPOINT ["/app/server"]
