FROM golang:1.23.2-alpine3.20 AS builder

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o rkeeper .

FROM alpine:latest
COPY --from=builder /app /app

WORKDIR /app

EXPOSE 8080

CMD ["/app/rkeeper"]
