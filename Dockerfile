FROM golang:1.19-alpine as builder

ENV GO111MODULE=on

WORKDIR /app
COPY . .

# RUN apk --no-cache add git alpine-sdk build-base gcc

RUN cd src; go get

RUN cd src; go build -o main main.go

CMD ["./src/main"]