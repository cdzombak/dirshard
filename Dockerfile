ARG BIN_NAME=dirshard
ARG BIN_VERSION=<unknown>

FROM golang:1 AS builder
ARG BIN_NAME
ARG BIN_VERSION
WORKDIR /src/dirshard
COPY . .
RUN go build -ldflags="-X main.version=${BIN_VERSION}" -o ./out/${BIN_NAME} .

FROM scratch
ARG BIN_NAME
COPY --from=builder /src/dirshard/out/${BIN_NAME} /usr/bin/dirshard
RUN /usr/bin/dirshard -version
ENTRYPOINT ["/usr/bin/dirshard"]
