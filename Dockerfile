FROM golang:1.15.3

ENV GOPATH $GOPATH:/go
ENV PATH $PATH:$GOPATH/bin

COPY go.mod go.sum /go/src/
ADD . /go/src/app
WORKDIR /go/src/app

RUN apt-get update
RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/beego/bee"
RUN go get "github.com/beego/beego/v2"

COPY ./entrypoint.sh /
RUN chmod 777 /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]