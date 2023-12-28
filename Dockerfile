# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

WORKDIR /app
COPY . ./
RUN go mod vendor
RUN go build -o /amongis

EXPOSE 8080

CMD [ "/amongis" ]