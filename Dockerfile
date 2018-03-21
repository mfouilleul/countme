FROM golang

COPY . /go/src/github.com/mfouilleul/countme

RUN go get github.com/go-redis/redis ;\
	go get net/http ;\
	go get github.com/ghodss/yaml ;\
	go get github.com/cenkalti/backoff;\
	go get github.com/sirupsen/logrus;

WORKDIR /go/src/github.com/mfouilleul/countme

RUN go build -o /go/bin/countme -ldflags "-X main.version=`cat VERSION`"

COPY config.yaml /etc/countme.yaml

ENTRYPOINT ["countme"]

CMD ["--config", "/etc/countme.yaml"]

EXPOSE 8000
