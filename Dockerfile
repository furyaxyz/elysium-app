# stage 1 Generate elysium-appd Binary
FROM docker.io/golang:1.20.4-alpine3.17 as builder
# hadolint ignore=DL3018
RUN apk update && apk add --no-cache \
    gcc \
    git \
    make \
    musl-dev
COPY . /elysium-app
WORKDIR /elysium-app
RUN make build

# stage 2
FROM docker.io/alpine:3.17.3

# Read here why UID 10001: https://github.com/hexops/dockerfile/blob/main/README.md#do-not-use-a-uid-below-10000
ARG UID=10001
ARG USER_NAME=elysium

ENV ELYSIUM_HOME=/home/${USER_NAME}

# hadolint ignore=DL3018
RUN apk update && apk add --no-cache \
    bash \
    # Creates a user with $UID and $GID=$UID
    && adduser ${USER_NAME} \
    -D \
    -g ${USER_NAME} \
    -h ${ELYSIUM_HOME} \
    -s /sbin/nologin \
    -u ${UID}

# Copy in the binary
COPY --from=builder /elysium-app/build/elysium-appd /bin/elysium-appd

COPY --chown=${USER_NAME}:${USER_NAME} docker/entrypoint.sh /opt/entrypoint.sh

USER ${USER_NAME}

# p2p, rpc and prometheus port
EXPOSE 26656 26657 1317 9090

ENTRYPOINT [ "/bin/bash", "/opt/entrypoint.sh" ]
