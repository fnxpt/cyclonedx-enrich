FROM golang:latest AS builder
WORKDIR /go/src/
ENV CGO_ENABLED=0

ADD . .

RUN GIN_MODE=release go build -ldflags="-s -w" -o release/cyclonedx-enrich .

FROM scratch AS runtime
WORKDIR /

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/release/cyclonedx-enrich .

ENTRYPOINT [ "/cyclonedx-enrich"]
