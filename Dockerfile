FROM golang:1.21.1

WORKDIR /usr/local/src

COPY ./ ./

RUN go mod tidy
RUN go build -o ./app_start ./cmd/pick_up_point_server/main.go

CMD ["./app_start"]
