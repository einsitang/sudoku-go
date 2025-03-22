# copyright (c) 2025 einsitang. All Rights Reserve.
#!/bin/bash

android_sdk=/Users/einsitang/Library/Android/sdk/ndk/29.0.13113456

platform="android"

# x86_64 for simulator
arch="amd64"
export GOOS="android"
export GOARCH="amd64"
export CGO_ENABLED="1"
export CC="$android_sdk/toolchains/llvm/prebuilt/darwin-x86_64/bin/x86_64-linux-android35-clang"
export CXX="$android_sdk/toolchains/llvm/prebuilt/darwin-x86_64/bin/x86_64-linux-android35-clang++"
go build -buildmode=c-shared -ldflags="-s -w" -o build/cgo/$platform/libsudoku_$arch.so ./cgo

# aarch64
arch="aarch64"
export GOOS="android"
export GOARCH="arm"
export CGO_ENABLED="1"
export CC="$android_sdk/toolchains/llvm/prebuilt/darwin-x86_64/bin/aarch64-linux-android35-clang"
export CXX="$android_sdk/toolchains/llvm/prebuilt/darwin-x86_64/bin/aarch64-linux-android35-clang++"
go build -buildmode=c-shared -ldflags="-s -w" -o build/cgo/$platform/libsudoku_$arch.so ./cgo

# armv7a
arch="armv7a"
export GOARCH="arm"
export CC="$android_sdk/toolchains/llvm/prebuilt/darwin-x86_64/bin/armv7a-linux-androideabi35-clang"
export CXX="$android_sdk/toolchains/llvm/prebuilt/darwin-x86_64/bin/armv7a-linux-androideabi35-clang++"
go build -buildmode=c-shared -ldflags="-s -w" -o build/cgo/$platform/libsudoku_$arch.so ./cgo


# use command nm to see detail ->  nm -g -D build/cgo/android/libsudoku_arm64.so | grep Gen