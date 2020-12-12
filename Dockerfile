FROM golang:1.15.3

COPY go.mod go.sum /go/src/
ADD . /go/src
WORKDIR /go/src

RUN apt-get update
RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/beego/bee"
RUN go get "github.com/astaxie/beego"

COPY ./entrypoint.sh /
RUN chmod 777 /entrypoint.sh

# ENTRYPOINT ["/entrypoint.sh"]