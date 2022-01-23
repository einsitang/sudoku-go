package main

import (
	"testing"

	sudoku "github.com/forfuns/sudoku-go/core"
)

func BenchmarkMain(t *testing.B) {
	puzzle := [81]int8{
		-1, -1, 8 /* */, 9, -1, 6 /* */, -1, -1, 5,
		-1, 4, 3, -1 /* */, -1, -1, -1 /* */, 2, -1,
		-1, -1, -1 /* */, -1, -1, -1, -1 /* */, -1, -1,

		-1, -1, 4 /* */, -1, -1, -1 /* */, 9, -1, -1,
		5, -1, -1 /* */, -1, 4, -1 /* */, 6, 8, -1,
		-1, -1, -1 /* */, 1, -1, -1 /* */, -1, -1, -1,

		2, -1, -1 /* */, -1, 8, -1 /* */, -1, 7, -1,
		-1, -1, -1 /* */, -1, 3, 4 /* */, 1, -1, -1,
		-1, 6, -1 /* */, -1, -1, 9 /* */, -1, -1, -1,
	}

	_sudoku := new(sudoku.Sudoku)
	err := _sudoku.Init(puzzle)
	if err != nil {
		t.Log(err)
		t.Fail()
	} else {
		_sudoku.Debug()
		// t.Log("sudo puzzle is : %v",_sudoku.Puzzle())
		// t.Log("sudo answer is : %v",_sudoku.Answer())
		t.Log("sudoku is done")
	}

}
