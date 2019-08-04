FROM golang:1.12.7-stretch AS build-env

RUN go env GOARCH > /goarch && \
    go env GOARM > /goarm && \
    go env GOOS > /goos

ENV GOPATH=/go
COPY . /go/src/github.com/grafana/loki
WORKDIR /go/src/github.com/grafana/loki
# TODO: This line should be implement with Makefile rules
RUN cd /go/src/github.com/grafana/loki/fluent-bit && \
  CGO_ENABLED=1 GOOS=$(cat /goos) GOARCH=$(cat /goarch) GOARM=$(cat /goarm) && \
  go build -buildmode=c-shared -o out_loki.so .

FROM fluent/fluent-bit:1.2
COPY --from=build-env /go/src/github.com/grafana/loki/fluent-bit/out_loki.so /usr/lib/x86_64-linux-gnu/
COPY fluent-bit/docker/fluent-bit-loki.conf \
     /fluent-bit/etc/
EXPOSE 2020

# Entry point
CMD ["/fluent-bit/bin/fluent-bit", "-c", "/fluent-bit/etc/fluent-bit-loki.conf", "-e", "/usr/lib/x86_64-linux-gnu/out_loki.so"]
