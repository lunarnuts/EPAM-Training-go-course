FROM golang:1.16-alpine

WORKDIR /storeapp

COPY util ./util
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-store-server

EXPOSE 8081

CMD [ "/docker-store-server" ]
