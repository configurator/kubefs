#!/usr/bin/env bash

set -euo pipefail -o posix
cd "$(dirname "$0")"

# List all possible platforms
# NOTE: Not all of these combinations have been tested - this are just a list of
# combinations where a build is possible
platforms=(
    darwin/386
    darwin/amd64
    freebsd/386
    freebsd/amd64
    freebsd/arm
    linux/386
    linux/amd64
    linux/arm
    linux/arm64
    linux/ppc64
    linux/ppc64le
    linux/mips
    linux/mipsle
    linux/mips64
    linux/mips64le
    linux/s390x
)

targetDir=dist

mkdir -p "$targetDir"

packageName="kubefs"
package="."

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    export GOOS=${platform_split[0]}
    export GOARCH=${platform_split[1]}
    out="$targetDir/$packageName-$GOOS-$GOARCH"
    if [ $GOOS = "windows" ]; then
        out+='.exe'
    fi

    echo "Building for $platform..."
    go build -o "$out" "$package"
    echo
done
