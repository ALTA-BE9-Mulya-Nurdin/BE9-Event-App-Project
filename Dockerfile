FROM golang:1.17

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go build -o alta-event

CMD ["./alta-event"]