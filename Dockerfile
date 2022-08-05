ARG GOLANG_VERSION=1.19
ARG GOLANG_OPTIONS="CGO_ENABLED=0 GOOS=linux GOARCH=amd64"

ARG PROJECT="gcp-status"

FROM golang:${GOLANG_VERSION} as build

ARG VERSION=""
ARG COMMIT=""

ARG PROJECT

WORKDIR /${PROJECT}

COPY go.* ./
COPY *.go .

RUN env ${GOLANG_OPTIONS} \
    go build \
    -ldflags "-X main.OSVersion=${VERSION} -X main.GitCommit=${COMMIT}" \
    -a -installsuffix cgo \
    -o /go/bin/${PROJECT} \
    .

FROM gcr.io/distroless/base-debian11

ARG PROJECT

LABEL org.opencontainers.image.source https://github.com/DazWilkin/gcp-status

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /go/bin/${PROJECT} /exporter

EXPOSE 8080

ENTRYPOINT ["/exporter"]
CMD ["--endpoint=:9989","--path=/metrics"]
