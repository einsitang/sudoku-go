package main

// #cgo CFLAGS: -g -Wall
// #include <stdio.h>
// #include <stdlib.h>
/*
struct sudoku_channel {
 int err;
 signed char* matrix;
};
*/
import "C"
import (
	"unsafe"

	"github.com/einsitang/sudoku-go/v2/internal/core"
	"github.com/einsitang/sudoku-go/v2/sudoku"
)

//export Generate
func Generate(level C.int, output unsafe.Pointer) {
	_level := int(level)
	_sudoku, _err := sudoku.Generate(_level)
	puzzle := _sudoku.Puzzle()

	// 将 buff 转换为 *[81]int8 类型
	// buf := (*[81]int8)(output)
	buf := (*C.struct_sudoku_channel)(output)
	if _err != nil {
		buf.err = 1
	} else {
		buf.err = 0
		buf.matrix = (*C.schar)(unsafe.Pointer(&puzzle[0]))
	}

	// 调用方完事了记得给我free output
}

//export Solve
func Solve(input unsafe.Pointer, isStrict C.int, output unsafe.Pointer) {
	isStrictBool := isStrict != 0
	puzzle := [81]int8{}
	inputPuzzle := (*[81]C.schar)(input) // 将 C 指针转换为 [81]int 类型

	for i := 0; i < 81; i++ {
		puzzle[i] = int8(inputPuzzle[i]) // 解引用并逐元素复制
	}
	var _sudoku core.Sudoku
	var _err error
	if isStrictBool {
		_sudoku, _err = sudoku.Solve(puzzle, sudoku.WithStrict())
	} else {
		_sudoku, _err = sudoku.Solve(puzzle)
	}

	// 将 buff 转换为 *[81]int8 类型
	// buf := (*[81]int8)(output)

	var solution [81]int8
	buf := (*C.struct_sudoku_channel)(output)
	if _err != nil {
		buf.err = 1
		return
	}
	buf.err = 0
	solution = _sudoku.Solution()

	buf.matrix = (*C.schar)(unsafe.Pointer(&solution[0]))

	// 调用方完事了记得给我 free input & output
}

func main() {}
