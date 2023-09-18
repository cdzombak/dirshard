ARG BIN_NAME
ARG BIN_VERSION

FROM golang:1 AS builder
WORKDIR /src/dirshard
COPY . .
RUN go build -ldflags="-X main.version=${BIN_VERSION}" -o ./out/${BIN_NAME} .

FROM scratch
COPY --from=builder /src/dirshard/out/${BIN_NAME} /usr/bin/dirshard
ENTRYPOINT ["/usr/bin/dirshard"]
