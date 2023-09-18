FROM golang:1 AS builder
RUN apt-get update \
    && DEBIAN_FRONTEND=noninteractive apt-get -y --no-install-recommends install -yq \
       build-essential \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /src/dirshard
COPY . .
RUN make build

FROM scratch
COPY --from=builder /src/dirshard/out/dirshard /usr/bin/dirshard
ENTRYPOINT ["/usr/bin/dirshard"]
