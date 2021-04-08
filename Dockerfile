FROM alpine:3.11 AS alpine

FROM golang:1.13 AS build-stage

ARG APIGATEWAY_BINARY=apigateway_binary
ARG SERVICEA_BINARY=servicea_binary
ARG SERVICEB_BINARY=serviceb_binary

WORKDIR /project

COPY . /project
RUN go mod download \
  # build api gateway
    && go build -o ${APIGATEWAY_BINARY} cmd/apigateway/* \
    # build service a
    && go build -o ${SERVICEA_BINARY} cmd/servicea/* \
    # build service b
    && go build -o ${SERVICEB_BINARY} cmd/serviceb/* 


FROM alpine AS apigateway
WORKDIR /app
COPY --from=build-stage /project/apigateway_binary /app/apigateway_binary
RUN apk add libc6-compat curl \
  && chmod +x /app/apigateway_binary

EXPOSE 80/tcp

ENTRYPOINT [ "/app/apigateway_binary" ]

FROM alpine AS servicea
WORKDIR /app
COPY --from=build-stage /project/servicea_binary /app/servicea_binary
RUN apk add libc6-compat curl \
  && chmod +x /app/servicea_binary

EXPOSE 80/tcp

ENTRYPOINT [ "/app/servicea_binary" ]


FROM alpine AS serviceb
WORKDIR /app
COPY --from=build-stage /project/serviceb_binary /app/serviceb_binary
RUN apk add libc6-compat curl \
  && chmod +x /app/serviceb_binary

EXPOSE 80/tcp

ENTRYPOINT [ "/app/serviceb_binary" ]