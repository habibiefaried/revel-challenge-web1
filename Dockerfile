FROM golang:1.17.0-alpine3.13 AS build-env

ENV CGO_ENABLED 0
ENV GOPATH=/go
ENV PATH=$PATH:$GOPATH/bin

RUN apk add --no-cache git
ADD . /app
WORKDIR /app

# Install revel framework
RUN go get -u github.com/revel/revel
RUN go get -u github.com/revel/cmd/revel

RUN revel build -m prod . /tmp/testbuild

FROM alpine:3.13.6
EXPOSE 9000

ARG PROOFTXT
ENV PROOFTXT $PROOFTXT

RUN echo "$PROOFTXT" > /root/proof.txt 2>/dev/null

COPY --from=build-env /tmp/testbuild /app
WORKDIR /app

RUN chmod +x run.sh

ENTRYPOINT ["/bin/sh"]
CMD ["-c", "./run.sh"]