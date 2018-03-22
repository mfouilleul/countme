FROM golang

COPY . /go/src/github.com/mfouilleul/countme
WORKDIR /go/src/github.com/mfouilleul/countme

RUN go get -u github.com/golang/dep/cmd/dep; \
    dep ensure

RUN go build -o /go/bin/countme -ldflags "-X main.version=`cat VERSION`"

COPY config.yaml /etc/countme.yaml

ENTRYPOINT ["countme"]

CMD ["--config", "/etc/countme.yaml"]

EXPOSE 8000
