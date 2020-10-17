FROM golang:alpine
RUN apk update && apk upgrade
WORKDIR /go
COPY code /go/code
COPY modules /go/modules
RUN sed -i '/^replace/d' /go/code/go.mod
RUN sed -i '/anshumanpatil/d' /go/code/go.mod
RUN sed -i '5 i replace github.com/anshumanpatil/goserver-print-api => /go/modules' /go/code/go.mod
WORKDIR /usr/local/go/src/github.com/anshumanpatil/goserver-print-api
COPY modules/. .
WORKDIR /go/code
RUN go clean --modcache
CMD ["go", "run", "main.go"]
