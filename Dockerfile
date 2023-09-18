FROM golang:1 AS builder
WORKDIR /src/dirshard
COPY . .
RUN make build

FROM scratch
COPY --from=builder /src/dirshard/out/dirshard /usr/bin/dirshard
ENTRYPOINT ["/usr/bin/dirshard"]
