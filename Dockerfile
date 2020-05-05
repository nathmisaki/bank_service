FROM golang:1.14.2-alpine
WORKDIR /go/src/app
RUN apk update && \
    apk upgrade 
RUN apk add git
RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm


ADD . .
CMD ["go", "run", "main.go"]