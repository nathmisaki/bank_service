FROM golang:1.14.2-alpine
WORKDIR /go/src/app
ADD . .
RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm


CMD ["go", "run", "main.go"]