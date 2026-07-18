FROM golang:1.26-alpine3.23 AS builder

COPY . /build/
RUN cd /build && ./build.sh

FROM alpine:3.23

RUN apk add --no-cache bash curl jq git ffmpeg \
	# Python for Google Chat
	python3 py3-pip py3-setuptools py3-wheel \
	# Common dependencies that need native extensions for Python bridges
	py3-magic py3-ruamel.yaml py3-aiohttp py3-pillow py3-olm py3-pycryptodome \
	# Lottieconverter for animated sticker conversion
	lottieconverter

# Necessary to self-compile bridges
RUN apk add --no-cache build-base olm-dev ffmpeg-libs python3-dev py3-virtualenv giflib rust cargo cmake protoc clang-dev zlib-dev

ENV GO_VERSION=1.26.3
RUN wget https://go.dev/dl/go${GO_VERSION}.linux-arm64.tar.gz && \
	rm -rf /usr/local/go && \
	tar -C /usr/local -xzf go${GO_VERSION}.linux-arm64.tar.gz && \
	rm go${GO_VERSION}.linux-arm64.tar.gz

ENV GOROOT=/usr/local/go
ENV PATH=/usr/local/go/bin:$PATH

COPY --from=builder /build/bbctl /usr/local/bin/bbctl
COPY ./docker/run-bridge.sh /usr/local/bin/run-bridge.sh
ENV SYSTEM_SITE_PACKAGES=true
VOLUME /data

ENTRYPOINT ["/usr/local/bin/run-bridge.sh"]
