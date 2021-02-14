FROM golang:1.15-alpine

RUN mkdir /app

COPY . /app

WORKDIR /app

EXPOSE 8080

RUN go build -o server .

CMD [ "/app/server" ]
