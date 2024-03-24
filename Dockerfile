FROM golang:latest

WORKDIR /app

COPY go.mod go.sum .env togos.db ./

RUN go mod download

COPY . .

RUN go build -o tg .

# EXPOSE 8080

CMD ["./tg"]