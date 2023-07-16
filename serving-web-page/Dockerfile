FROM golang:1.18-alpine

WORKDIR /serving-web-page

COPY go.mod ./
COPY *.go ./
COPY static ./static

RUN go build -o /main

EXPOSE 8080

CMD [ "/main" ]