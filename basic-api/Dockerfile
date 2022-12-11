FROM golang:1.1-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o /basic-api

EXPOSE 9190

CMD [ "/basic-api" ]