FROM golang:1.7.3-alpine
RUN apk add --no-cache bash git
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . /go/src/app
RUN go get github.com/lib/pq
RUN go-wrapper download
RUN go-wrapper install
CMD ["go-wrapper", "run"]
