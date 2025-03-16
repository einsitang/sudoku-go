#!/bin/bash

# 编译真机 (arm64)
export GOOS=ios
export GOARCH=arm64
export CGO_ENABLED=1
export CC=$(xcrun --sdk iphoneos --find clang)
export CGO_CFLAGS="-arch arm64 -isysroot $(xcrun --sdk iphoneos --show-sdk-path) -fembed-bitcode"
# export CGO_LDFLAGS="-arch arm64 -isysroot $(xcrun --sdk iphoneos --show-sdk-path) -Wl,-install_name,@rpath/libexample.dylib"
go build -buildmode=c-archive -o build/cgo/ios/libsudoku_arm64.a ./cgo

# 编译模拟器 (x86_64 / arm64)
export GOARCH=amd64
# export GOARCH=arm64
export CGO_ENABLED=1
export CC=$(xcrun --sdk iphonesimulator --find clang)
export CGO_CFLAGS="-arch x86_64 -isysroot $(xcrun --sdk iphonesimulator --show-sdk-path)"
# export CGO_LDFLAGS="-arch arm64 -isysroot $(xcrun --sdk iphonesimulator --show-sdk-path) -Wl,-install_name,@rpath/libexample.dylib"
go build -buildmode=c-archive -o build/cgo/ios/libsudoku_simulator_$GOARCH.a ./cgo

# 使用 xcodebuild 组合成xcframework
# xcodebuild -create-xcframework -library build/cgo/ios/libsudoku_x86_64.a -library build/cgo/ios/libsudoku_arm64.a -output ./build/cgo/ios/libsudoku.xcframework

# 使用lipo组合静态库
# 组合 
# lipo -create -output build/cgo/ios/libsudoku.a build/cgo/ios/libsudoku_x86_64.a build/cgo/ios/libsudoku_arm64.a

# 查看
# lipo -info build/cgo/ios/libsudoku.a

# 静态库 转 动态库 方式 <- 未验证

# 解压静态库，得出一堆 *.o 文件
# ar -x libsudoku.a 
# 重新组合成动态库
# xcrun -sdk iphoneos clang -arch arm64 -fpic -shared -Wl,-all_load *.o -framework CoreFoundation -o libsudoku.dylib -isysroot $(xcrun --sdk iphoneos --show-sdk-path)
# 验证
# file libsudoku.dylib -> libsudoku.dylib: Mach-O 64-bit dynamically linked shared library arm64
