#!/usr/bin/env bash
set -euo pipefail

[ -z "$(git tag --points-at HEAD)" ] && echo "$(git describe --always --long --dirty | sed 's/^v//')" || echo "$(git tag --points-at HEAD | sed 's/^v//')"
