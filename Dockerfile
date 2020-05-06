FROM golang:1.14.2-alpine
WORKDIR /go/src/app
RUN apk update && \
    apk upgrade 
RUN apk add git
RUN apk add build-base
RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN go get -u github.com/onsi/ginkgo/ginkgo
RUN go get -u github.com/onsi/gomega/... 


ADD . .
CMD ["go", "run", "main.go"]