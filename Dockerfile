FROM golang:1.14.2-alpine
RUN apk update && \
    apk upgrade 
RUN apk add git
RUN apk add build-base
RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN go get -u github.com/onsi/ginkgo/ginkgo
RUN go get -u github.com/onsi/gomega/... 
RUN go get github.com/githubnemo/CompileDaemon

WORKDIR /go/src/github.com/nelsonmhjr/bank_service
COPY . .
ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o api main.go" -command="./api"
