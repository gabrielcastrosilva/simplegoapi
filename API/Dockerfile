FROM golang:1.9.2
ADD . /go/src/api
WORKDIR /go/src/api
RUN go get github.com/urfave/negroni
RUN go install
EXPOSE 3000
ENTRYPOINT ["/go/bin/api"]
