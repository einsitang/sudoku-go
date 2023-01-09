package main

import (
	sudoku "github.com/einsitang/sudoku-go/core"
	"testing"
)

//func BenchmarkMain(t *testing.B) {
//	t.Logf("Benchmark N : %v", t.N)
//	err, _sudoku := getSudoku(false)
//	if err != nil {
//		t.Log(err)
//		t.Fail()
//	} else {
//		_sudoku.Debug()
//		t.Log("sudoku is done")
//	}
//}

func BenchmarkNormalMode(t *testing.B) {
	t.Logf("BenchmarkNormalMode N : %v", t.N)
	err, _ := getSudoku(false)
	if err != nil {
		t.Log(err)
		t.Fail()
	} else {
		//_sudoku.Debug()
		t.Log("sudoku is done")
	}
}

func BenchmarkStrictInit(t *testing.B) {
	t.Logf("BenchmarkStrictInit N : %v", t.N)
	err, _ := getSudoku(true)
	if err != nil {
		t.Log(err)
		t.Fail()
	} else {
		//_sudoku.Debug()
		t.Log("sudoku is done")
	}
}

func TestSolve(t *testing.T) {
	err, _sudoku := getSudoku(true)
	if err != nil {
		t.Log(err)
		t.Fail()
	} else {
		_sudoku.Debug()
		t.Log("sudoku is done")
	}
}

func getSudoku(isStrictMode bool) (err error, _sudoku sudoku.Sudoku) {
	//puzzle := [81]int8{
	//	-1, -1, 8 /* */, 9, -1, 6 /* */, -1, -1, 5,
	//	-1, 4, 3, -1 /* */, -1, -1, -1 /* */, 2, -1,
	//	-1, -1, -1 /* */, -1, -1, -1, -1 /* */, -1, -1,
	//
	//	-1, -1, 4 /* */, -1, -1, -1 /* */, 9, -1, -1,
	//	5, -1, -1 /* */, -1, 4, -1 /* */, 6, 8, -1,
	//	-1, -1, -1 /* */, 1, -1, -1 /* */, -1, -1, -1,
	//
	//	2, -1, -1 /* */, -1, 8, -1 /* */, -1, 7, -1,
	//	-1, -1, -1 /* */, -1, 3, 4 /* */, 1, -1, -1,
	//	-1, 6, -1 /* */, -1, -1, 9 /* */, -1, -1, -1,
	//}

	// this puzzle from generator with level expert
	puzzle := [81]int8{9, 8, -1, -1, -1, -1, 7, -1, 6, -1, -1, -1, 7, -1, 4, -1, -1, 2, -1, -1, -1, 9, -1, -1, -1, -1, 5, -1, -1, -1, 1, 8, -1, -1, -1, -1, -1, -1, 4, -1, -1, -1, 1, 9, -1, -1, -1, -1, 5, -1, -1, -1, -1, -1, -1, -1, 9, -1, -1, -1, -1, -1, 1, 4, -1, -1, -1, 5, -1, -1, -1, -1, 6, 7, -1, -1, -1, 3, 2, -1, -1}

	_sudoku = sudoku.Sudoku{}
	if isStrictMode {
		err = _sudoku.StrictInit(puzzle)
	} else {
		err = _sudoku.Init(puzzle)
	}

	return
}
