ARG GOLANG_VERSION=1.20

ARG PROJECT="gcp-status"

FROM docker.io/golang:${GOLANG_VERSION} as build

ARG VERSION=""
ARG COMMIT=""

ARG PROJECT

WORKDIR /${PROJECT}

COPY go.* ./
COPY *.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -ldflags "-X main.OSVersion=${VERSION} -X main.GitCommit=${COMMIT}" \
    -a -installsuffix cgo \
    -o /go/bin/${PROJECT} \
    .

FROM gcr.io/distroless/static

ARG PROJECT

LABEL org.opencontainers.image.source https://github.com/DazWilkin/gcp-status

COPY --from=build /go/bin/${PROJECT} /exporter

EXPOSE 8080

ENTRYPOINT ["/exporter"]
CMD ["--endpoint=:9989","--path=/metrics"]
