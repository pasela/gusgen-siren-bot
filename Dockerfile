FROM golang:alpine AS builder

RUN apk update && apk upgrade && \
    apk add --no-cache g++ git

COPY go.mod go.sum /app/gusgen-siren-bot/
RUN cd /app/gusgen-siren-bot && go mod download

COPY . /app/gusgen-siren-bot
WORKDIR /app/gusgen-siren-bot
RUN go build -a -ldflags="-s -w"

#----------------------------------------

FROM golang:alpine
LABEL maintainer="Yuki <paselan@gmail.com>"

# zoneinfo and certificates
# ENV ZONEINFO=/zoneinfo.zip
# ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=builder /app/gusgen-siren-bot/gusgen-siren-bot /gusgen-siren-bot

WORKDIR /

ENTRYPOINT ["/gusgen-siren-bot"]
