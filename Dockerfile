FROM golang:1.18.0-alpine3.15 as builder

LABEL author="blackfurystation"

RUN apk add --update-cache \
    git \
    gcc \
    musl-dev \
    linux-headers \
    make \
    wget

RUN git clone https://github.com/blackfurystation/blackfury.git /blackfury && \
    #chmod -R 755 /blackfury && \
    chmod -R 755 /blackfury
WORKDIR /blackfury
RUN make install

# final image
FROM golang:1.18.0-alpine3.15

RUN mkdir -p /data

VOLUME ["/data"]

COPY --from=builder /go/bin/blackfuryd /usr/local/bin/blackfuryd

EXPOSE 26656 26657 1317 9090

ENTRYPOINT ["blackfuryd"]