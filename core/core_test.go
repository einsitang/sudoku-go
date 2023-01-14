package core

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestSolver(t *testing.T) {
	hard := "8..........36......7..9.2...5...7.......457.....1...3...1....68..85...1..9....4.."
	puzzle := [81]int8{}

	for i, char := range hard {
		if char == '.' {
			puzzle[i] = -1
		} else {
			puzzle[i] = (int8)(char - '0')
		}
	}
	beginTime := time.Now()
	_sudoku := Sudoku{}
	_sudoku.Init(puzzle)
	endTime := time.Now()
	t.Logf("total time : %v'ms", endTime.Sub(beginTime).Milliseconds())
}

func BenchmarkInit(t *testing.B) {
	t.Logf("BenchmarkInit N : %v", t.N)
	beginTime := time.Now()
	for n := 0; n < t.N; n++ {
		err, _ := getSudoku(false)
		if err != nil {
			t.Log(err)
			t.Fail()
		} else {
			//_sudoku.Debug()
			//t.Log("sudoku is done")
		}
	}
	endTime := time.Now()
	t.Logf("BenchmarkInit (N:%v) total time : %v ms", t.N, endTime.Sub(beginTime).Milliseconds())
}

func BenchmarkStrictInit(t *testing.B) {
	t.Logf("BenchmarkStrictInit N : %v", t.N)
	beginTime := time.Now()
	for n := 0; n < t.N; n++ {
		err, _ := getSudoku(true)
		if err != nil {
			t.Log(err)
			t.Fail()
		} else {
			//_sudoku.Debug()
			//t.Log("sudoku is done")
		}
	}
	endTime := time.Now()
	t.Logf("BenchmarkStrictInit (N:%v) total time : %v ms", t.N, endTime.Sub(beginTime).Milliseconds())
}

func TestSolve(t *testing.T) {
	err, _sudoku := getSudoku(false)
	if err != nil {
		t.Log(err)
		t.Fail()
	} else {
		_sudoku.Debug()
		t.Log("sudoku is done")
	}
}

func getSudoku(isStrictMode bool) (err error, _sudoku Sudoku) {
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
	//puzzle := [81]int8{9, 8, -1, -1, -1, -1, 7, -1, 6, -1, -1, -1, 7, -1, 4, -1, -1, 2, -1, -1, -1, 9, -1, -1, -1, -1, 5, -1, -1, -1, 1, 8, -1, -1, -1, -1, -1, -1, 4, -1, -1, -1, 1, 9, -1, -1, -1, -1, 5, -1, -1, -1, -1, -1, -1, -1, 9, -1, -1, -1, -1, -1, 1, 4, -1, -1, -1, 5, -1, -1, -1, -1, 6, 7, -1, -1, -1, 3, 2, -1, -1}
	puzzle := [81]int8{-1, -1, -1, -1, 8, 4, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, -1, -1, 1, 6, 9, -1, -1, -1, -1, -1, -1, -1, -1, 5, 7, -1, -1, -1, 9, -1, 8, -1, -1, -1, 9, -1, 3, -1, -1, -1, -1, 1, -1, -1, 6, -1, -1, -1, -1, -1, 8, -1, 3, -1, -1, 4, -1, -1, -1, 7, 9, 2, -1, -1, -1, -1, -1, -1, 3, -1, -1, -1, 6, -1, 5}
	strs := []string{}
	for _, num := range puzzle {
		if num == -1 {
			strs = append(strs, ".")
			continue
		}
		strs = append(strs, strconv.Itoa((int)(num)))
	}
	fmt.Printf("%v\n", strings.Join(strs, ""))
	_sudoku = Sudoku{}
	if isStrictMode {
		err = _sudoku.StrictInit(puzzle)
	} else {
		err = _sudoku.Init(puzzle)
	}

	return
}
