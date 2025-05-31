#!/usr/bin/env sh

root=`git rev-parse --show-toplevel`
if [ "$root" != `pwd` ]; then
	echo "Not at the root of the repo - exiting"
	exit 1
fi

for d in $(find ./pkg -type d -mindepth 1)
do
	echo "golangci-lint fmt $d"
	golangci-lint fmt $d
	echo "golangci-lint run --fix $d"
	golangci-lint run --fix $d
done

echo "building package generator"
mkdir -p bin
for arch in arm64 amd64; do
	for os in linux darwin; do
		mkdir -p bin/linux/${os}/${arch}
		echo "building for ${os}/${arch}"
		env GOOS=${os} GOARCH=${arch} go build -o "./bin/${os}/${arch}/generator" generator/generator.go
	done
done

echo "building new linetypes"
go run ./generator -version=v12.0 > pkg/linetypes/v12_0/linetypes.go \
  && go run ./generator -version=v11.1 > pkg/linetypes/v11_1/linetypes.go \
  && go run ./generator -version=v11.0 > pkg/linetypes/v11_0/linetypes.go \
  && go run ./generator -version=v10.1 > pkg/linetypes/v10_1/linetypes.go \
  && go run ./generator -version=v10.0 > pkg/linetypes/v10_0/linetypes.go \
  && golangci-lint fmt ./...

echo "format and lint linetypes"
echo "formatting ./pkg/linetypes"
for f in $(find ./pkg -name linetypes.go)
do
	echo "golangci-lint fmt $f"
	golangci-lint fmt $f
	echo "golangci-lint run --fix $f"
	golangci-lint run --fix $f
done

echo "building sample program"
for arch in arm64 amd64; do
	for os in linux darwin; do
		echo "building for ${os}/${arch}"
		env GOOS=${os} GOARCH=${arch} go build -o "./bin/${os}/${arch}/sample_parser" examples/sample_parser/sample_parser.go
	done
done

echo "finished"
exit 0
