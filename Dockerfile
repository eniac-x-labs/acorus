FROM --platform=$BUILDPLATFORM golang:1.21.1-alpine3.18 as builder

RUN apk add --no-cache make ca-certificates gcc musl-dev linux-headers git jq bash

COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum

WORKDIR /app

RUN go mod download

# build acorus with the shared go.mod & go.sum files
COPY ./acorus /app/acorus
COPY event/processors/op-stack/mantle/op-bindings /app/op-bindings
COPY synchronizer/op-service /app/op-service
COPY synchronizer/op-node /app/op-node

WORKDIR /app/acorus

RUN make acorus

FROM alpine:3.18

COPY --from=builder /app/acorus/acorus /usr/local/bin
COPY --from=builder /app/acorus/acorus.yaml /app/acorus/acorus.yaml
COPY --from=builder /app/acorus/migrations /app/acorus/migrations

WORKDIR /app

ENV INDEXER_MIGRATIONS_DIR="/app/acorus/migrations"

CMD ["acorus", "index", "--config", "/app/acorus/acorus.yaml"]
