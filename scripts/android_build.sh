# copyright (c) 2025 einsitang. All Rights Reserve.
#!/bin/bash

android_sdk=/Users/einsitang/Library/Android/sdk/ndk/25.2.9519653

platform="android"


export GOOS="android"
export CGO_ENABLED="1"
# aarch64
arch="aarch64"
export GOARCH="arm64"
export CC="$android_sdk/toolchains/llvm/prebuilt/darwin-x86_64/bin/aarch64-linux-android33-clang"
# export CXX="$android_sdk/toolchains/llvm/prebuilt/darwin-x86_64/bin/aarch64-linux-android33-clang++"
export CGO_CFLAGS="-DLIBSUDOKU_VERSION=\"1.0.0\""
go build -buildmode=c-shared -ldflags="-s -w" -o build/cgo/$platform//arm64-v8a/libsudoku.so ./cgo

# armv7a
arch="armv7a"
export GOARCH="arm"
export CC="$android_sdk/toolchains/llvm/prebuilt/darwin-x86_64/bin/armv7a-linux-androideabi33-clang"
# export CXX="$android_sdk/toolchains/llvm/prebuilt/darwin-x86_64/bin/armv7a-linux-androideabi33-clang++"
export CGO_CFLAGS="-DLIBSUDOKU_VERSION=\"1.0.0\""
go build -buildmode=c-shared -ldflags="-s -w" -o build/cgo/$platform/armeabi-v7a/libsudoku.so ./cgo


# use command nm to see detail ->  nm -g -D build/cgo/android/libsudoku_arm64.so | grep Gen