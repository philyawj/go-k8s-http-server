FROM golang:alpine

WORKDIR /app
COPY . /app

ENV BG_COLOR=pink
ENV HEADING_COLOR=steelblue
ENV HEADING_TEXT='Docker Golang HTTP Server'

RUN go install

CMD ["/go/bin/go-k8s-http-server"]