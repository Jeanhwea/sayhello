FROM golang:1.20

WORKDIR /app
ADD . /app
RUN go mod tidy && go build -o server

EXPOSE 8080

ENTRYPOINT ["/app/server"]
