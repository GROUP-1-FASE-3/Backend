FROM golang:1.19

# membuat direktori app
RUN mkdir /app

# set working directory /app
WORKDIR /app

COPY ./ /app

RUN go mod tidy

RUN go build -o GP3K1-api

CMD ["./GP3K1-api"]