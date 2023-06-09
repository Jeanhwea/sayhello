FROM golang:1.20-alpine as builder

WORKDIR /app
ADD . /app
RUN go env -w GOPROXY=https://goproxy.cn,direct && \
    go build -o server

FROM golang:1.20-alpine as runner
WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 80

ENTRYPOINT ["/app/server"]
