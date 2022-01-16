FROM golang

ENV GOBIN=/go/bin/ 
ENV GO111MODULE=on 
ENV GOPROXY=https://goproxy.cn

COPY server.go /go/src

WORKDIR /go
RUN go build -o bin/server src/server.go

ENTRYPOINT ["bin/server"]
