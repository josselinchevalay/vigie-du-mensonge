FROM golang:1.25.0-alpine

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /docker-gs-ping

CMD [ "/docker-gs-ping" ]