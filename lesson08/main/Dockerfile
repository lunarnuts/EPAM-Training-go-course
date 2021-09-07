FROM golang:1.16-alpine

WORKDIR /serverapp

COPY util ./util
COPY go.mod ./
COPY index.html ./
#COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-main-server

EXPOSE 8080

CMD [ "/docker-main-server" ]
