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

echo "building package buildHeaderLineTypes"
mkdir -p bin
for arch in arm64 amd64; do
	for os in linux darwin; do
		mkdir -p bin/linux/${os}/${arch}
		echo "building for ${os}/${arch}"
		env GOOS=${os} GOARCH=${arch} go build -o "./bin/${os}/${arch}/buildHeaderLineTypes" pkg/buildHeaderLineTypes/buildHeaderLineTypes.go
	done
done

echo "building new metLineTypeDefinitions.go"
go run pkg/buildHeaderLineTypes/buildHeaderLineTypes.go -version=v12.0 > pkg/metLineTypeDefinitions_v12_0/metLineTypeDefinitions.go
go run pkg/buildHeaderLineTypes/buildHeaderLineTypes.go -version=v11.1 > pkg/metLineTypeDefinitions_v11_1/metLineTypeDefinitions.go
go run pkg/buildHeaderLineTypes/buildHeaderLineTypes.go -version=v11.0 > pkg/metLineTypeDefinitions_v11_0/metLineTypeDefinitions.go
go run pkg/buildHeaderLineTypes/buildHeaderLineTypes.go -version=v10.1 > pkg/metLineTypeDefinitions_v10_1/metLineTypeDefinitions.go
go run pkg/buildHeaderLineTypes/buildHeaderLineTypes.go -version=v10.0 > pkg/metLineTypeDefinitions_v10_0/metLineTypeDefinitions.go

echo "format and lint metLinetypeDefinitions"
echo "formatting ./pkg/metLineTypeDefinitions"
for f in $(find ./pkg -name metLineTypeDefinitions.go)
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
		env GOOS=${os} GOARCH=${arch} go build -o "./bin/${os}/${arch}/sample_parser" pkg/sample_parser/sample_parser.go
	done
done

echo "finished"
exit 0
