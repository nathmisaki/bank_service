FROM golang:1.14.2-alpine
RUN apk update && \
    apk upgrade 
RUN apk add git
RUN apk add build-base
WORKDIR /go/src/github.com/nelsonmhjr/bank_service
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go get -u github.com/onsi/ginkgo/ginkgo
RUN go get github.com/githubnemo/CompileDaemon
CMD CompileDaemon -log-prefix=false -build="go build -o api main.go" -command="./api"
