ARG GOLANG_VERSION=1.23.3

ARG PROJECT="gcp-status"

ARG TARGETOS
ARG TARGETARCH

ARG VERSION
ARG COMMIT

FROM --platform=${TARGETARCH} docker.io/golang:${GOLANG_VERSION} AS build

ARG PROJECT

WORKDIR /${PROJECT}

COPY go.* ./
COPY *.go .

ARG TARGETOS
ARG TARGETARCH

ARG VERSION
ARG COMMIT

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build \
    -ldflags "-X main.OSVersion=${VERSION} -X main.GitCommit=${COMMIT}" \
    -a -installsuffix cgo \
    -o /go/bin/${PROJECT} \
    .

FROM --platform=${TARGETARCH} gcr.io/distroless/debian12-static:latest

ARG PROJECT

LABEL org.opencontainers.image.source=https://github.com/DazWilkin/gcp-status

COPY --from=build /go/bin/${PROJECT} /exporter

EXPOSE 8080

ENTRYPOINT ["/exporter"]
CMD ["--endpoint=:9989","--path=/metrics"]
