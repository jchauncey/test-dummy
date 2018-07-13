FROM quay.io/deis/go-dev:v1.12.2

ADD . /go/src/github.com/jchauncey/test-dummy
WORKDIR /go/src/github.com/jchauncey/test-dummy

CMD go run main.go