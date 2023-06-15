# stage 1 Generate elysium-appd Binary
FROM golang:1.18-alpine as builder
# hadolint ignore=DL3018
RUN apk update && apk --no-cache add make gcc musl-dev git
COPY . /elysium-app
WORKDIR /elysium-app
RUN make build

# stage 2
FROM alpine:3.18.2
# hadolint ignore=DL3018
RUN apk update && apk --no-cache add bash

COPY --from=builder /elysium-app/build/elysium-appd /bin/elysium-appd
COPY  docker/entrypoint.sh /opt/entrypoint.sh

# p2p, rpc and prometheus port
EXPOSE 26656 26657 1317 9090

ENV ELYSIUM_HOME /opt

ENTRYPOINT [ "/bin/bash", "/opt/entrypoint.sh" ]
