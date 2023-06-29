package core

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

func puzzles() [][81]int8 {
	return [][81]int8{
		// easy
		{-1, -1, 5, 4, -1, 7, 3, -1, 1, 4, -1, 7, 3, 1, -1, 6, -1, -1, -1, 8, 1, 2, 6, -1, -1, 9, -1, 5, -1, -1, -1, 3, 9, -1, -1, -1, -1, 2, 6, 7, -1, 4, -1, -1, -1, 7, -1, 3, 5, -1, -1, 9, -1, -1, 6, -1, 2, -1, -1, -1, 8, -1, 9, -1, 3, -1, -1, 7, 2, 4, 5, -1, 8, 7, 4, -1, -1, -1, 1, 3, 2},
		{-1, 1, -1, 8, -1, -1, -1, -1, 3, -1, 7, -1, -1, 9, 3, -1, -1, -1, 2, 9, -1, -1, 5, -1, -1, -1, -1, 1, -1, -1, -1, -1, 9, 3, 5, -1, 5, -1, 4, 7, 1, -1, -1, 9, 2, 9, -1, 2, 5, 3, -1, 1, -1, 8, 7, -1, -1, 6, 8, -1, -1, -1, -1, 6, 2, 9, 3, 4, -1, 5, 8, -1, 3, -1, 8, -1, 2, -1, 7, 4, 6},
		// expert
		{-1, 7, -1, -1, 2, 6, 1, -1, -1, -1, -1, -1, 9, -1, -1, -1, -1, -1, -1, -1, 8, 7, -1, -1, -1, -1, -1, -1, 9, -1, 4, 8, -1, -1, -1, -1, 1, 2, -1, -1, -1, 7, -1, 3, -1, -1, -1, -1, 3, -1, 2, 4, -1, -1, -1, 3, 9, 2, 5, -1, -1, 8, -1, 4, -1, -1, -1, 6, -1, -1, 7, -1, 6, 5, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, 7, -1, -1, -1, -1, -1, -1, -1, -1, -1, 6, 5, 2, -1, -1, -1, 1, -1, 2, -1, -1, -1, 3, -1, -1, 9, -1, -1, -1, -1, -1, 7, 4, 5, 4, -1, -1, -1, -1, 5, 8, -1, -1, -1, 2, -1, 7, -1, 3, -1, -1, 9, -1, -1, -1, 1, -1, -1, -1, -1, 8, -1, -1, 3, 2, -1, -1, 9, -1, -1, -1, 8, 9, -1, -1, 7, 1, -1, -1},
		{-1, -1, 1, -1, 8, 9, -1, -1, -1, 2, 8, -1, -1, -1, 7, -1, -1, 1, -1, 7, -1, 3, -1, 1, -1, 9, -1, 5, -1, -1, -1, 9, -1, 6, 1, 7, -1, -1, -1, -1, -1, -1, -1, 3, -1, -1, 6, 7, 5, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, 4, -1, 2, -1, 7, -1, -1, -1, -1, 5, -1, -1, -1, -1, 6, -1, 9, -1, -1},
		// hell
		{-1, -1, 1, -1, -1, -1, 4, 2, -1, -1, -1, 7, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 8, -1, -1, -1, -1, 5, -1, -1, 4, -1, 9, -1, -1, -1, -1, -1, -1, -1, 5, -1, -1, 8, 3, -1, 9, -1, -1, -1, -1, -1, 7, -1, 7, -1, -1, -1, 3, -1, 2, -1, 5, -1, 3, -1, -1, 9, 8, -1, -1, -1, -1, -1, -1, -1, -1, -1, 1, -1, -1},
		{-1, -1, -1, 8, 5, -1, 7, -1, -1, -1, -1, 4, -1, 3, -1, -1, -1, -1, -1, 2, 3, -1, -1, -1, -1, -1, 4, 1, -1, -1, -1, -1, -1, -1, -1, -1, 9, -1, -1, -1, 1, 7, -1, -1, -1, -1, -1, -1, 4, -1, -1, -1, -1, 6, -1, 1, -1, -1, -1, -1, -1, 7, -1, -1, -1, -1, 3, 9, -1, 1, 8, -1, -1, -1, -1, 5, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, 7, -1, -1, -1, -1, 4, -1, 3, -1, -1, -1, -1, -1, 2, 3, -1, -1, -1, -1, -1, 4, 1, -1, -1, -1, -1, -1, -1, -1, -1, 9, -1, -1, -1, 1, 7, -1, -1, -1, -1, -1, -1, 4, -1, -1, -1, -1, 6, -1, 1, -1, -1, -1, -1, -1, 7, -1, -1, -1, -1, 3, 9, -1, 1, 8, -1, -1, -1, -1, 5, -1, -1, -1, -1, -1},
		{-1, -1, -1, 2, 6, 5, 1, -1, -1, -1, 6, -1, -1, -1, -1, -1, 7, 4, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 3, -1, 5, -1, 9, -1, -1, 8, -1, 4, -1, 6, -1, -1, 9, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 7, 4, -1, -1, -1, -1, -1, 5, 1, -1, -1, -1, -1, -1, 8, -1, 3, -1, -1, -1, -1, -1, -1, -1, 7},
	}
}

func dfsSolver(puzzle [81]int8, t *testing.T) {
	beginTime := time.Now()
	_sudoku := Sudoku{}
	err := _sudoku.Init(puzzle)
	if err != nil {
		t.Logf("❌[dfs] ..")
	}
	endTime := time.Now()
	subTime := endTime.Sub(beginTime)
	t.Logf("solution : %v", sudokuGoPuzzle2str(&_sudoku.answer))
	t.Logf("total time : %v'ms | %v'ns [dfs]			with puzzle \n%v", subTime.Milliseconds(), subTime.Nanoseconds(), puzzle)
}

func dfsWithStrictSolver(puzzle [81]int8, t *testing.T) {
	beginTime := time.Now()
	_sudoku := Sudoku{}
	err := _sudoku.StrictInit(puzzle)
	if err != nil {
		t.Logf("❌[dfs-Strict] ..")
	}
	endTime := time.Now()
	subTime := endTime.Sub(beginTime)
	t.Logf("solution : %v", sudokuGoPuzzle2str(&_sudoku.answer))
	t.Logf("total time : %v'ms | %v'ns [dfs-Strict]		with puzzle \n%v", subTime.Milliseconds(), subTime.Nanoseconds(), puzzle)
}

func dlxSolver(puzzle [81]int8, t *testing.T) {
	// puzzleStr := puzzle2str(puzzle)
	beginTime := time.Now()
	solution := DLXSolve(puzzle)
	endTime := time.Now()
	subTime := endTime.Sub(beginTime)
	t.Logf("solution : %v", solution)
	t.Logf("total time : %v'ms | %v'ns [dxl]			with puzzle \n%v", subTime.Milliseconds(), subTime.Nanoseconds(), puzzle)
}

func TestSolvers(t *testing.T) {
	puzzles := puzzles()
	for _, puzzle := range puzzles {
		dfsSolver(puzzle, t)
		dfsWithStrictSolver(puzzle, t)
		dlxSolver(puzzle, t)
		t.Log("================================================")
	}
}

func TestSolveHellLevelPuzzle(t *testing.T) {
	t.Log("try to solve puzzle with hell level use [dfs] and [dlx]")
	// hard := "8..........36......7..9.2...5...7.......457.....1...3...1....68..85...1..9....4.."
	// puzzle := Str2sudokuGo(&hard)
	hellPuzzle := [81]int8{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 3, -1, -1, 4, 5, -1, 9, 7, 1, -1, -1, -1, -1, -1, 6, 1, -1, -1, 6, -1, 7, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 3, -1, 6, -1, -1, -1, -1, -1, -1, 9, -1, -1, -1, -1, -1, -1, 1, 9, -1, -1, 2, -1, 3, -1, 4, -1, -1, -1, -1, 7, -1, -1, -1, 5, 8, -1, -1, -1}
	// hellPuzzleHard := SudokuGo2str(&hellPuzzle)
	beginTime := time.Now()
	_sudoku := Sudoku{}
	_sudoku.Init(hellPuzzle)
	endTime := time.Now()
	t.Logf("dfs total time : %v'ms [dfs(maybe not!!!)]", endTime.Sub(beginTime).Milliseconds())

	beginTime = time.Now()
	_sudoku = Sudoku{}
	solution := _sudoku.DLXInit(hellPuzzle)
	t.Logf("solution : %v", solution)
	endTime = time.Now()
	t.Logf("dlx total time : %v'ms [dlx]", endTime.Sub(beginTime).Milliseconds())
}

func BenchmarkInit(t *testing.B) {
	t.Logf("BenchmarkInit N : %v", t.N)
	beginTime := time.Now()
	for n := 0; n < t.N; n++ {
		_, err := getSudoku(false)
		if err != nil {
			t.Log(err)
			t.Fail()
		}
	}
	endTime := time.Now()
	t.Logf("BenchmarkInit (N:%v) total time : %v ms", t.N, endTime.Sub(beginTime).Milliseconds())
}

func BenchmarkStrictInit(t *testing.B) {
	t.Logf("BenchmarkStrictInit N : %v", t.N)
	beginTime := time.Now()
	for n := 0; n < t.N; n++ {
		_, err := getSudoku(true)
		if err != nil {
			t.Log(err)
			t.Fail()
		}
	}
	endTime := time.Now()
	t.Logf("BenchmarkStrictInit (N:%v) total time : %v ms", t.N, endTime.Sub(beginTime).Milliseconds())
}

func TestSolve(t *testing.T) {
	_sudoku, err := getSudoku(false)
	if err != nil {
		t.Log(err)
		t.Fail()
	} else {
		_sudoku.Debug()
		t.Log("sudoku is done")
	}
}

func getSudoku(isStrictMode bool) (_sudoku Sudoku, err error) {
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
