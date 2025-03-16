# copyright (c) 2025 einsitang. All Rights Reserve.
# for einsitang macos env
#!/bin/bash

go build -ldflags="-s -w" -buildmode=c-shared -o build/cgo/libsudoku.so ./cgo