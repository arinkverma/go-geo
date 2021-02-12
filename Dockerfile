FROM golang:1.12-alpine

RUN apk update && apk add bash git
RUN go get github.com/gin-gonic/gin
RUN go get github.com/gomodule/redigo/redis

WORKDIR /go/src/app
COPY src .
COPY ./tmp/cities15000.txt .

#RUN go get -d -v ./...
RUN go install -v ./...
RUN chmod +x /go/bin/app
CMD ["/go/bin/app"]
EXPOSE 8080:8080

#CMD ["/bin/sh"]
