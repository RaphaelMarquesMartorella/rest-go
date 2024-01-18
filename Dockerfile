FROM golang:1.20.13

WORKDIR /usr/src/app

COPY . .

RUN go build -o main ./cmd/server

CMD [ "./main" ]