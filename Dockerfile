FROM golang:1.17.2
COPY ./src /go/src/github.com/eturnhout/live-shapes
WORKDIR /go/src/github.com/eturnhout/live-shapes
RUN go build .
EXPOSE 8080
CMD ["./go-live-shapes"]
