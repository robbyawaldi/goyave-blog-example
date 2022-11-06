FROM golang:1.17

LABEL maintainer="Jérémy LAMBERT (SystemGlitch) <jeremy.la@outlook.fr>"

RUN go get github.com/cespare/reflex

RUN wget https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-alpine-linux-amd64-v0.6.1.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-v0.6.1.tar.gz \
    && rm dockerize-alpine-linux-amd64-v0.6.1.tar.gz

WORKDIR /app

EXPOSE 8080

CMD dockerize -wait tcp://mariadb:3306 reflex -s -- sh -c 'go run main.go'