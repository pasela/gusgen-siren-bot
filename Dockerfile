FROM golang:1.19 AS builder

RUN apt-get update && apt-get install -y \
    git \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -a -ldflags="-s -w" -trimpath

#----------------------------------------

FROM golang:1.19
LABEL maintainer="Yuki <paselan@gmail.com>"

COPY --from=builder /usr/src/app/gusgen-siren-bot /gusgen-siren-bot

WORKDIR /

ENTRYPOINT ["/gusgen-siren-bot"]
