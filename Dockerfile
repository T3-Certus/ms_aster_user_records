FROM golang:1.19.4-alpine

RUN mkdir -p /app
WORKDIR /app

COPY . /app

RUN go build -o main .

EXPOSE 3000-3001

CMD ["./main"]