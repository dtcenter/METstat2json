#!/usr/bin/env sh

root=`git rev-parse --show-toplevel`
if [ "$root" != `pwd` ]; then
	echo "Not at the root of the repo - exiting"
	exit 1
fi

echo "formatting primary packages"
gofumpt -w ./pkg/buildHeaderLineTypeUtilities ./pkg/buildHeaderLineTypes ./pkg/structColumnDefs ./pkg/structColumnTypes ./pkg/sample_parser

echo "linting primary packages"
golangci-lint run --fix ./pkg/buildHeaderLineTypeUtilities/... ./pkg/buildHeaderLineTypes/... ./pkg/structColumnDefs/... ./pkg/structColumnTypes/... ./pkg/sample_parser/...

echo "building package buildHeaderLineTypes"
go build -o ./bin/ pkg/buildHeaderLineTypes/buildHeaderLineTypes.go

echo "saving current structColumnTypes.go to /tmp/structColumnTypes.go.bak"
mv pkg/structColumnTypes/structColumnTypes.go pkg/structColumnTypes/structColumnTypes.go.bak

echo "building new structColumnTypes.go"
./bin/buildHeaderLineTypes > pkg/structColumnTypes/structColumnTypes.go

echo "format and lint structColumnTypes"
echo "formatting ./pkg/structColumnTypes"
gofumpt -w ./pkg/structColumnTypes
echo "linting ./pkg/structColumnTypes"
golangci-lint run --fix ./pkg/structColumnTypes/...

echo "building sample program"
go build -o ./bin/sample_parser pkg/sample_parser/sample_parser.go
echo "finished"
exit 0