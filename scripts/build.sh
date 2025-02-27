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
mkdir -p bin
for arch in arm64 amd64; do
	for os in linux darwin; do
		mkdir -p bin/linux/${os}/${arch}
		echo "building for ${os}/${arch}"
		env GOOS=${os} GOARCH=${arch} go build -o "./bin/${os}/${arch}/buildHeaderLineTypes" pkg/buildHeaderLineTypes/buildHeaderLineTypes.go
	done
done

echo "saving current structColumnTypes.go to /tmp/structColumnTypes.go.bak"
mv pkg/structColumnTypes/structColumnTypes.go pkg/structColumnTypes/structColumnTypes.go.bak

echo "building new structColumnTypes.go"
go run pkg/buildHeaderLineTypes/buildHeaderLineTypes.go > pkg/structColumnTypes/structColumnTypes.go

echo "format and lint structColumnTypes"
echo "formatting ./pkg/structColumnTypes"
gofumpt -w ./pkg/structColumnTypes
echo "linting ./pkg/structColumnTypes"
golangci-lint run --fix ./pkg/structColumnTypes/...

echo "building sample program"
for arch in arm64 amd64; do
	for os in linux darwin; do
		echo "building for ${os}/${arch}"
		env GOOS=${os} GOARCH=${arch} go build -o "./bin/${os}/${arch}/sample_parser" pkg/sample_parser/sample_parser.go
	done
done

echo "finished"
exit 0