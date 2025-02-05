FROM golang:1.23.4-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN cd cmd/gavialis && go build

EXPOSE 8000

CMD [ "./cmd/gavialis/gavialis" ]