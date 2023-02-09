FROM golang:1.16-alpine

WORKDIR /abhishek

COPY ./src/ /abhishek/

RUN go build -o /service-broker

CMD ["/service-broker"]