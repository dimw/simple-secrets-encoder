FROM golang:1.16-alpine

WORKDIR /go/src/app
COPY . /go/src/app

RUN go get -d -v
RUN go install -v
RUN go build -o /usr/local/bin/sse

WORKDIR /workspace

ENTRYPOINT ["sse"]
CMD ["help"]
