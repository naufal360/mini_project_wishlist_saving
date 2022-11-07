FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /docker-wishlist

EXPOSE 80

CMD [ "/docker-wishlist" ]