FROM golang:1.22.1 AS builder

SHELL ["/bin/bash", "-c"]

COPY . /go/random-person-service

WORKDIR /go/random-person-service

RUN set -x && CGO_ENABLED=0 go build .


FROM scratch

COPY --from=builder /go/random-person-service/random-person-service /random-person-service

CMD ["/random-person-service"]