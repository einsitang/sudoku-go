# copyright (c) 2025 einsitang. All Rights Reserve.
#!/bin/bash

platform="ios"
arch="arm64"

export GOOS="ios"
export GOARCH="arm64"
export CGO_ENABLED="1"
export CGO_CFLAGS="-fembed-bitcode"

# SDK 
# xcodebuild -showsdks

# ios arm
export CC=$GOROOT/misc/ios/clangwrap.sh GOOS=$GOOS GOARCH=$GOARCH SDK=iphoneos PLATFORM=ios -v
go build -ldflags="-s -w" -buildmode=c-archive -o build/cgo/$platform/libsudoku_$arch.a ./cgo

# x86 模拟器
export CC=$GOROOT/misc/ios/clangwrap.sh GOOS=$GOOS GOARCH=amd64 SDK=iphonesimulator PLATFORM=ios-simulator
go build -ldflags="-s -w" -buildmode=c-archive -o build/cgo/$platform/libsudoku_x86_64.a ./cgo

# lipo -create build/cgo/ios/libsudoku_arm64.a build/cgo/ios/libsudoku_x86_64.a -output build/cgo/libsudoku_ios.a

#组合
# lipo -create out_ios/libevparser.a out_x86/libevparser.a -output libevparser.a

# 查看便衣结果
# lipo -info libevparser.a