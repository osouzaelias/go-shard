FROM golang:1.21
LABEL maintainer="osouzaelias@gmail.com"

COPY . /app

WORKDIR /app/cmd

RUN go build -o go-shard

CMD ["./go-shard"]