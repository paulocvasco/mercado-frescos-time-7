FROM golang:1.18-bullseye

WORKDIR ./meli

COPY go.mod .

COPY . .

RUN go mod tidy -v
RUN go build ./cmd/server/

EXPOSE 8080

CMD [ "./server" ]