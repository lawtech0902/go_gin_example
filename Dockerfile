FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/lawtech0902/go_gin_example
COPY . $GOPATH/src/github.com/lawtech0902/go_gin_example
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go_gin_example"]