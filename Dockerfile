# syntax = docker/dockerfile:1

FROM golang:1.16-alpine 

WORKDIR /random-api-go


COPY go.mod ./
COPY go.sum ./
RUN go mod download 


COPY *.go ./
COPY ./random/*.go ./random/

RUN go build -o /random-api-go

EXPOSE 8080


CMD [ "./random-api-go" ]