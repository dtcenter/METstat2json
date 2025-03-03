#!/usr/bin/env sh

root=`git rev-parse --show-toplevel`
if [ "$root" != `pwd` ]; then
	echo "Not at the root of the repo - exiting"
	exit 1
fi
echo "formatting"
gofumpt -w ./pkg/buildHeaderLineTypeUtilities ./pkg/buildHeaderLineTypes ./pkg/structColumnDefs ./pkg/structColumnTypes
echo "linting"
golangci-lint run --fix ./pkg/buildHeaderLineTypeUtilities/... ./pkg/buildHeaderLineTypes/... ./pkg/structColumnDefs/... ./pkg/structColumnTypes/...
echo "finished"
exit 0
