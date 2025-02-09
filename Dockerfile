# Builder
FROM golang:1.22-alpine AS builder
RUN apk add --update make git curl

ARG MODULE_NAME=homework

COPY Makefile /home/${MODULE_NAME}/Makefile
COPY go.mod /home/${MODULE_NAME}/go.mod
COPY go.sum /home/${MODULE_NAME}/go.sum

WORKDIR /home/${MODULE_NAME}

COPY . /home/${MODULE_NAME}

RUN make build

# Service
FROM alpine:latest as server
ARG MODULE_NAME=homework
WORKDIR /root/

COPY --from=builder /home/${MODULE_NAME}/main .
COPY --from=builder /home/${MODULE_NAME}/server.crt .
COPY --from=builder /home/${MODULE_NAME}/server.key .

RUN chown root:root main

CMD ["./main"]
