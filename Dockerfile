# build
FROM            golang:1.27-alpine as builder
RUN             apk add --no-cache git gcc musl-dev make
ENV             GO111MODULE=on
WORKDIR         /go/src/moul.io/fs-bundler
COPY            go.* ./
RUN             go mod download
COPY            . ./
RUN             make install

# minimalist runtime
FROM            alpine:3.24
COPY            --from=builder /go/bin/fs-bundler /bin/
ENTRYPOINT      ["/bin/fs-bundler"]
CMD             []
