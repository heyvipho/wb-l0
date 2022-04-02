FROM golang:1.18.0-alpine3.15

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./
RUN go install -v ./

EXPOSE 80
CMD ["order-sender"]