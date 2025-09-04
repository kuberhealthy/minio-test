FROM golang:1.24 AS builder
WORKDIR /build
COPY go.mod go.sum /build/

COPY . /build
WORKDIR /build/cmd/minio-test
ENV CGO_ENABLED=0
RUN go build -v
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/cmd/minio-test/minio-test /app/minio-test
ENTRYPOINT ["/app/minio-test"]
