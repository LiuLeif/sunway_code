FROM golang:1.18.2-buster as builder
ENV GO111MODULE=off
COPY . /app
ENV GOPATH=/app
RUN GOOS=linux GOARCH=amd64 go install hexintek.com/license_server

FROM debian:buster
COPY --from=builder /app/bin /app/bin
COPY --from=builder /app/src/hexintek.com/license_server /app/src/hexintek.com/license_server
WORKDIR /app
CMD ./bin/license_server