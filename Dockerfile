# Build Stage
FROM golang:1.12.6-stretch AS build

WORKDIR /go/src/github.com/samuelkadolph/ifconfig

COPY . .

RUN CGO_ENABLED=0 make

# Run Stage
FROM alpine:latest

WORKDIR /root

COPY --from=build /go/src/github.com/samuelkadolph/ifconfig/bin/ifconfig .

EXPOSE 7570

CMD ["./ifconfig"]
