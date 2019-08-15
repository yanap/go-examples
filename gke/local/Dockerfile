FROM golang:1.12
WORKDIR /go/src/github.com/yanap/go-examples
ENV GO111MODULE=on
RUN go get github.com/pilu/fresh
EXPOSE 8080
CMD ["fresh"]