FROM golang:1.14.2-alpine
WORKDIR /go/src/app
ADD . .
RUN go get github.com/gin-gonic/gin


CMD ["go", "run", "main.go"]