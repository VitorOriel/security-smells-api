FROM golang:1.22 AS builder

ADD security-smells-api/ app/
WORKDIR app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/smelly-kube-api

FROM alpine:latest AS final
COPY --from=builder /bin/smelly-kube-api /bin/smelly-kube-api

ENTRYPOINT ["/bin/smelly-kube-api"]