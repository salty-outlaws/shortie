FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache bash git make

WORKDIR /app
COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -tags netgo -a -v -installsuffix cgo -o bin/shortie main.go

FROM alpine:3
RUN apk update \
    && apk add --no-cache curl wget \
    && apk add --no-cache ca-certificates \
    && update-ca-certificates 2>/dev/null || true

COPY --from=builder /app/bin/shortie /shortie

CMD ["/shortie"]

ENV PORT=8000
EXPOSE $PORT
