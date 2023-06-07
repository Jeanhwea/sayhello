FROM golang:1.20

CMD go mod tidy && go build -o server

EXPOSE 8080

ENTRYPOINT ["./server"]
