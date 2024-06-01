FROM golang:1.22-alpine

ENV TZ=Asia/Jakarta

WORKDIR /app

RUN apk add --update curl && \
    rm -rf /var/cache/apk/*

RUN go install github.com/cosmtrek/air@latest

RUN mkdir -p "/etc/air"

COPY .docker/config/air-conf.toml /etc/air/air-conf.toml

RUN sed -i "s/service_name/dbo-rest/g" "/etc/air/air-conf.toml"

ENTRYPOINT ["air", "-c", "/etc/air/air-conf.toml"]
