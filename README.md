# dirshard

Produce sharded path fragments from a filename.

## Usage
```
      dirshard [OPTIONS] -- some_object_key
 (or) cat object_list.txt | dirshard [OPTIONS]
```

Produces a path fragment consisting of the first N alphanumeric characters of the given object key, separated by a path separator.
No leading or trailing slash is produced.

### Options

- `-ci`: Case-insensitive: letters will be converted to lowercase. Can be set by the DIRSHARD_CI environment variable; this flag overrides the env var.
- `-n int`: Number of shards to produce. Can be set by the DIRSHARD_N environment variable; this flag overrides the env var. (default 1)
- `-skip`: Skip non-alphanumeric characters entirely, rather than converting them to underscores. Can be set by the DIRSHARD_SKIP environment variable; this flag overrides the env var.
- `-version`: Print version and exit.

### Examples

```
$ dirshard abcdef.txt
a

$ dirshard -n 3 abcdef.txt
a/b/c

$ dirshard -skip -n 2 a%^bcd.txt
a/b

$ dirshard -n 3 -ci Ab.txt
a/b/_
```

### Read from standard input

`dirshard` can read content (eg. a list of filenames) from standard input, outputting newline-delimited path fragments:

```
$ echo -e "Foo.txt\nBar.txt" | dirshard -ci
f
b
```

## Installation

### macOS via Homebrew

```shell
brew install cdzombak/oss/dirshard
```

### Debian via PackageCloud

Install my PackageCloud Debian repository if you haven't already:
```shell
curl -s https://packagecloud.io/install/repositories/cdzombak/oss/script.deb.sh?any=true | sudo bash
```

Then install `dirshard` via `apt-get`:
```shell
sudo apt-get install dirshard
```

### Manual installation from build artifacts

Pre-built binaries for Linux and macOS on various architectures are downloadable from each [GitHub Release](https://github.com/cdzombak/dirshard/releases). Debian packages for each release are available as well. 

### Build and install locally

```shell
git clone https://github.com/cdzombak/dirshard.git
cd dirshard
make build

cp out/dirshard $INSTALL_DIR
```

## Docker images

Docker images are available for a variety of Linux architectures from [Docker Hub](https://hub.docker.com/r/cdzombak/dirshard) and [GHCR](https://github.com/cdzombak/dirshard/pkgs/container/dirshard). Images are based on the `scratch` image and are as small as possible.

Run them via, for example:

```shell
docker run --rm cdzombak/dirshard:1 -n 2 abcdef.txt
docker run --rm ghcr.io/cdzombak/dirshard:1 -ci FOOBAR.txt
```

## About

- Issues: https://github.com/cdzombak/dirshard/issues/new
- Author: [Chris Dzombak](https://www.dzombak.com)
    - [GitHub: @cdzombak](https://www.github.com/cdzombak)

## License

LGPLv3; see `LICENSE` in this repository.
