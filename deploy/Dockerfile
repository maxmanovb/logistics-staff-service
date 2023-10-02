# Stage #1
FROM golang:1.21 AS builder

COPY ./ /go/src/app/

WORKDIR /go/src/app

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app ./main.go

# Stage #2
FROM alpine:latest AS base

RUN apk --no-cache add ca-certificates
RUN apk add --update tzdata

COPY --from=builder /go/bin/app /go/bin/app

# Stage #3
FROM scratch

COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /go/bin/app /go/bin/app


EXPOSE 3030

ENTRYPOINT ["/go/bin/app"]