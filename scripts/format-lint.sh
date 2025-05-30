#!/usr/bin/env sh

root=$(git rev-parse --show-toplevel)
if [ "$root" != $(pwd) ]; then
	echo "Not at the root of the repo - exiting"
	exit 1
fi
echo "formatting"
golangci-lint fmt ./...
echo "linting"
golangci-lint run --fix ./...
echo "finished"
exit 0
