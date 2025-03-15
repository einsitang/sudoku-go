package sudoku

import (
	"testing"
	"time"

	core "github.com/einsitang/sudoku-go/v2/internal/core"
	generator "github.com/einsitang/sudoku-go/v2/internal/generator"
)

func DLX() {
	// beginTime := time.Now()
	_hellSudoku, _ := generator.Generate(generator.LEVEL_HELL)
	_ = _hellSudoku
	// endTime := time.Now()
	// hellPuzzle := _hellSudoku.Puzzle()
	// t.Logf("hell puzzle :%v", hellPuzzle)
	// t.Logf("generate time :%v'ms ", endTime.Sub(beginTime).Milliseconds())

	// beginTime = time.Now()
	// _sudoku, err := core.Solve(hellPuzzle, &core.SudokuOption{IsOneSolutionMode: true})
	// endTime = time.Now()
	// t.Logf("err : %v", err)
	// t.Logf("solution : %v", _sudoku.Solution())
	// t.Logf("dfs total time : %v'ms [dfs with strict]", endTime.Sub(beginTime).Milliseconds())

	// beginTime = time.Now()
	// _sudoku, err = core.Solve(hellPuzzle, &core.SudokuOption{DLXMode: true})
	// endTime = time.Now()
	// t.Logf("err : %v", err)
	// t.Logf("solution : %v", _sudoku.Solution())
	// t.Logf("dlx total time : %v'ms [dlx]", endTime.Sub(beginTime).Milliseconds())

	// str, solved := core.DLXSolve2(hellPuzzle)
	// dlxSolution := core.Str2sudokuGo(&str)
	// t.Log("=======")
	// t.Logf("solved : %v", solved)
	// t.Logf("solution : %v", dlxSolution)

}

func TestDLXSolveSpeed(t *testing.T) {
	beginTime := time.Now()
	_hellSudoku, _ := generator.Generate(generator.LEVEL_HELL)
	_ = _hellSudoku
	endTime := time.Now()
	hellPuzzle := _hellSudoku.Puzzle()
	t.Logf("puzzle   : %v", hellPuzzle)
	t.Logf("generate time :%v'ms ", endTime.Sub(beginTime).Milliseconds())

	beginTime = time.Now()
	_sudoku, err := core.Solve(hellPuzzle, &core.SudokuOption{})
	endTime = time.Now()
	t.Logf("err : %v", err)
	t.Logf("solution : %v", _sudoku.Solution())
	t.Logf("dfs total time : %v'ns [dfs]", endTime.Sub(beginTime).Nanoseconds())

	beginTime = time.Now()
	_sudoku, err = core.Solve(hellPuzzle, &core.SudokuOption{DLXMode: true})
	endTime = time.Now()
	t.Logf("err : %v", err)
	t.Logf("solution : %v", _sudoku.Solution())
	t.Logf("dlx total time : %v'ns [dlx]", endTime.Sub(beginTime).Nanoseconds())

	beginTime = time.Now()
	_sudoku, err = core.Solve(hellPuzzle, &core.SudokuOption{IsOneSolutionMode: true})
	endTime = time.Now()
	t.Logf("err : %v", err)
	t.Logf("solution : %v", _sudoku.Solution())
	t.Logf("dfs total time : %v'ns [dfs with strict]", endTime.Sub(beginTime).Nanoseconds())

	beginTime = time.Now()
	_sudoku, err = Solve(hellPuzzle, WithStrict())
	endTime = time.Now()
	t.Logf("err : %v", err)
	t.Logf("solution : %v", _sudoku.Solution())
	t.Logf("dlx total time : %v'ns [auto with strict]", endTime.Sub(beginTime).Nanoseconds())

	beginTime = time.Now()
	solutionStr, solved := core.DLXStrictSolve(hellPuzzle)
	endTime = time.Now()
	t.Logf("solved : %v", solved)
	t.Logf("solution : %v", core.Str2sudokuGo(&solutionStr))
	t.Logf("dlx total time : %v'ns [dlx with strict]", endTime.Sub(beginTime).Nanoseconds())

	// str, solved := core.DLXSolve2(hellPuzzle)
	// dlxSolution := core.Str2sudokuGo(&str)
	// t.Log("=======")
	// t.Logf("solved : %v", solved)
	// t.Logf("solution : %v", dlxSolution)
}

func TestDLX(t *testing.T) {
	beginTime := time.Now()
	DLX()
	endTime := time.Now()
	t.Logf("generate time :%v'ms ", endTime.Sub(beginTime).Milliseconds())
}

func BenchmarkGeneratorWithDlx(b *testing.B) {
	DLX()
}
