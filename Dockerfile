FROM golang:1.14.9-alpine AS builder

WORKDIR /app

COPY . ./
RUN go mod download

RUN go build -o /mreader_bin

CMD [ "/mreader_bin" ]