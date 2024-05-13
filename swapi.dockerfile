FROM golang:1.20 AS build 

RUN mkdir /app

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 go build -o swapiApp  /app .


FROM alpine

RUN mkdir /app

WORKDIR /app

COPY --from=build /app/swapiApp  /app

CMD [ "/app/swapiApp" ]