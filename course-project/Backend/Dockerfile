FROM golang:1.16-alpine

WORKDIR /serverapp

COPY src ./src
WORKDIR /serverapp/src
RUN go mod download

WORKDIR /serverapp/src/cmd/rest-api
RUN go build -o /docker-main-server
EXPOSE 8080

CMD [ "/docker-main-server" ]