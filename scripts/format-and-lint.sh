#!/bin/sh

set -e

SCRIPT_DIR=$(dirname -- "$0")
SCRIPT_DIR=$(cd -- "$SCRIPT_DIR" && pwd)

"$SCRIPT_DIR/tools.sh" format-and-lint
