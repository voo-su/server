FROM alpine:3.20 AS builder

ARG GOLANG_VERSION=1.23.2

RUN apk update && \
    apk add --no-cache make gcc openssh bash musl-dev openssl-dev ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

RUN wget https://go.dev/dl/go$GOLANG_VERSION.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go$GOLANG_VERSION.linux-amd64.tar.gz && \
    rm go$GOLANG_VERSION.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin

RUN mkdir /usr/src

RUN mkdir /usr/src/voo-su

WORKDIR /usr/src/voo-su

COPY . ./

RUN make install

RUN make build

FROM alpine:3.20

COPY --from=builder /usr/src/voo-su/build /usr/bin

RUN mkdir /etc/voo-su

EXPOSE 8000 8001

CMD ["sh"]
