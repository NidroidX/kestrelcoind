ARG kestrelcoinD_IMAGE
ARG kestrelcoinMINER_IMAGE

FROM ${kestrelcoinD_IMAGE} as kestrelcoind
FROM ${kestrelcoinMINER_IMAGE} as kestrelcoinminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/NidroidX/kestrelcoind

WORKDIR /go/src/github.com/NidroidX/kestrelcoind

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=kestrelcoind /app/ /app/
COPY --from=kestrelcoinminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/NidroidX/kestrelcoind/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]