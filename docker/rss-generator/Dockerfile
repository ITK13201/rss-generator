ARG GOLANG_VERSION=1.22-bookworm

# Tools
FROM golang:${GOLANG_VERSION} AS tools

ARG GOOSE_VERSION=3.19.2
ARG AIR_VERSION=1.51.0
ARG COBRA_CLI_VERSION=1.3.0
ARG DELVE_VERSION=1.22.1
ARG DOCKERIZE_VERSION=0.7.0

RUN wget -O /usr/local/bin/goose https://github.com/pressly/goose/releases/download/v${GOOSE_VERSION}/goose_linux_x86_64 \
    && chmod +x /usr/local/bin/goose

RUN wget -O /usr/local/bin/air https://github.com/cosmtrek/air/releases/download/v${AIR_VERSION}/air_${AIR_VERSION}_linux_amd64 \
    && chmod +x /usr/local/bin/air

RUN go install github.com/spf13/cobra-cli@v${COBRA_CLI_VERSION} \
    && mv /go/bin/cobra-cli /usr/local/bin/cobra-cli

RUN go install github.com/go-delve/delve/cmd/dlv@v${DELVE_VERSION} \
    && mv /go/bin/dlv /usr/local/bin/dlv

RUN wget https://github.com/jwilder/dockerize/releases/download/v${DOCKERIZE_VERSION}/dockerize-linux-amd64-v${DOCKERIZE_VERSION}.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-v${DOCKERIZE_VERSION}.tar.gz \
    && rm -f dockerize-linux-amd64-v${DOCKERIZE_VERSION}.tar.gz


# Development environment
FROM golang:${GOLANG_VERSION} AS development

ENV GO111MODULE=on
ARG APP=rss-generator

ARG WORKDIR=/usr/local/src/${APP}
WORKDIR ${WORKDIR}

RUN apt-get update -y

COPY ./${APP}/go.mod .
COPY ./${APP}/go.sum .
RUN go mod download

COPY --from=tools /usr/local/bin/goose /usr/local/bin/goose
COPY --from=tools /usr/local/bin/air /usr/local/bin/air
COPY --from=tools /usr/local/bin/cobra-cli /usr/local/bin/cobra-cli
COPY --from=tools /usr/local/bin/dlv /usr/local/bin/dlv
COPY --from=tools /usr/local/bin/dockerize /usr/local/bin/dockerize

COPY ./${APP} .
COPY ./docker/${APP}/bin/docker-entrypoint.sh /usr/local/bin/docker-entrypoint.sh
COPY ./docker/${APP}/bin/goose.sh /usr/local/bin/goose.sh
RUN chmod +x /usr/local/bin/docker-entrypoint.sh \
    && chmod +x /usr/local/bin/goose.sh

EXPOSE 80

ENTRYPOINT ["sh", "/usr/local/bin/docker-entrypoint.sh"]
CMD ["-d", "-w", "-m"]


# Builder
FROM golang:${GOLANG_VERSION} AS builder

ENV GO111MODULE=on
ARG APP=rss-generator

ARG WORKDIR=/usr/local/src/${APP}
WORKDIR ${WORKDIR}

COPY ./${APP}/go.mod .
COPY ./${APP}/go.sum .
RUN go mod download

COPY ./${APP} .

RUN go build -o /usr/local/bin/${APP}


# Production environment
FROM gcr.io/distroless/base:debug AS production

ARG APP=rss-generator
WORKDIR /usr/local/src/${APP}

COPY --from=tools /usr/local/bin/goose /usr/local/bin/goose
COPY --from=tools /usr/local/bin/dockerize /usr/local/bin/dockerize
COPY --from=builder /usr/local/bin/${APP} /usr/local/bin/${APP}

COPY --chmod=775 ./docker/${APP}/bin/docker-entrypoint.sh /usr/local/bin/docker-entrypoint.sh
COPY --chmod=775 ./docker/${APP}/bin/goose.sh /usr/local/bin/goose.sh

COPY ./rss-generator/migrations ./migrations

ENTRYPOINT ["sh", "/usr/local/bin/docker-entrypoint.sh"]
CMD ["-w", "-m"]
