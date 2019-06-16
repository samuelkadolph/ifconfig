# Build Stage
FROM golang:1.12.6-stretch AS build

WORKDIR /go/src/github.com/samuelkadolph/ifconfig

COPY . .

RUN CGO_ENABLED=0 make

# Run Stage
FROM alpine:latest

MAINTAINER samuel@kadolph.com

ARG BUILD_DATE
ARG VCS_REF

LABEL org.label-schema.build-date=$BUILD_DATE
LABEL org.label-schema.name="ifconfig"
LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.vcs-ref=$VCS_REF
LABEL org.label-schema.vcs-url="https://github.com/samuelkadolph/ifconfig"

WORKDIR /root

COPY --from=build /go/src/github.com/samuelkadolph/ifconfig/bin/ifconfig .

EXPOSE 7570

CMD ["./ifconfig"]
