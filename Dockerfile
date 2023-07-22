FROM golang:1.20-alpine

WORKDIR /app

COPY ./ ./

RUN go mod download

RUN go build -o ./ ./cmd/travel_api/

EXPOSE 8080

CMD [ "./travel_api" ]