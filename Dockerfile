FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./

RUN go build -o ./bin/pricefetcher

EXPOSE 3030

CMD ["./bin/pricefetcher"]

