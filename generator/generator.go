package generator

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"time"

	sudoku "github.com/einsitang/sudoku-go/core"
)

const (
	LEVEL_EASY   = 0
	LEVEL_MEDIUM = 1
	LEVEL_HARD   = 2
	LEVEL_EXPERT = 3
	// LEVEL_HELL
	// ⚠️ this difficulty will take a long time , carefully to use
	LEVEL_HELL = 4

	// MIN_CONCURRENCY
	// minimum concurrency value is set here , if use concurrency to improve performance
	MIN_CONCURRENCY = 2

	// EMPTY puzzle empty value
	EMPTY = -1
)

var (
	// generate job limit times
	maxJobCount = 50 * runtime.NumCPU()
)

// Generate
// this function will generate sudoku with one-solution
func Generate(level int) (_sudoku sudoku.Sudoku, err error) {
	// concurrent generate
	// if level below medium , just use one concurrent that will work fine
	n := runtime.NumCPU()
	n = n>>2 + 1
	if n < MIN_CONCURRENCY {
		n = MIN_CONCURRENCY
	}

	digHoleTotal := sudoku.CONST_EASY_HOLES
	switch level {
	case LEVEL_EASY:
		digHoleTotal = sudoku.CONST_EASY_HOLES
		n = 1
	case LEVEL_MEDIUM:
		digHoleTotal = sudoku.CONST_MEDIUM_HOLES
		n = 1
	case LEVEL_HARD:
		digHoleTotal = sudoku.CONST_HARD_HOLES
	case LEVEL_EXPERT:
		digHoleTotal = sudoku.CONST_EXPERT_HOLES
	case LEVEL_HELL:
		digHoleTotal = sudoku.CONST_HELL_HOLES
		fmt.Printf("use concurrent : %v for \"LEVEL_HELL\" \n", n)
		fmt.Printf("😈 welcome to hell 😈 : this difficulty will take a long time...\n")
	default:
		err = errors.New("unknown level , make sure range by [0,3]")
		return
	}

	return doGenerate(digHoleTotal, n)
}

func doGenerate(digHoleTotal int, concurrency int) (_sudoku sudoku.Sudoku, err error) {

	sudokuCh := make(chan sudoku.Sudoku)
	// signal channel to make sure other goroutine will not block
	signal := make(chan int)
	done := false
	for i := 0; i < concurrency; i++ {
		go generate(sudokuCh, signal, digHoleTotal, &done, 1)
	}
	signal <- 1
	_sudoku = <-sudokuCh
	return
}

func generate(sudokuCh chan<- sudoku.Sudoku, signal chan int, digHoleTotal int, done *bool, jobCount int) {
	if *done {
		return
	}

	if jobCount >= maxJobCount {
		// reduce the difficulty
		oldDigHoleTotal := digHoleTotal
		digHoleTotal -= 2
		fmt.Printf("generate times : %d / %d(MAX) reduce the difficulty %d -> %d \n", jobCount, maxJobCount, oldDigHoleTotal, digHoleTotal)
		generate(sudokuCh, signal, digHoleTotal, done, 1)
		return
	}

	var simplePuzzle [81]int8

	// init simple puzzle
	nums := sudoku.ShuffleNumbers()
	ni := 0
	for i := range simplePuzzle {
		_, _, zone := sudoku.Location(i)
		simplePuzzle[i] = EMPTY

		// choose center zone to random fill
		if zone == 4 {
			simplePuzzle[i] = (int8)(nums[ni] + 1)
			ni++
		} else {
			simplePuzzle[i] = EMPTY
		}
	}

	basicSudoku := sudoku.Sudoku{}
	_ = basicSudoku.Init(simplePuzzle)

	// the dig hold process been pull away from function generate
	// because I wan't test each dig hole logic may faster
	// but only thing useful logic is try more times , now is twice => maxDigHoleProcessTimes := 2
	var resultSudoku *sudoku.Sudoku
	maxDigHoleProcessTimes := 2
	for resultSudoku == nil && maxDigHoleProcessTimes > 0 {
		resultSudoku = digHoleProcess(basicSudoku, digHoleTotal)
		maxDigHoleProcessTimes--
		if resultSudoku != nil {
			break
		}
	}

	if *done {
		// the work is done , don't need to check and send channel
		return
	}

	if resultSudoku == nil {
		// add job counter
		jobCount++
		generate(sudokuCh, signal, digHoleTotal, done, jobCount)
		return
	}

	*done = true
	doneAndCloseChannel(resultSudoku, signal, sudokuCh)

}

func doneAndCloseChannel(resultSudoku *sudoku.Sudoku, signal chan int, sudokuCh chan<- sudoku.Sudoku) {
	_, signalIsOpen := <-signal
	if signalIsOpen {
		sudokuCh <- *resultSudoku
		close(signal)
		close(sudokuCh)
	}
}

// dig hole process logic
func digHoleProcess(basicSudoku sudoku.Sudoku, digHoleTotal int) *sudoku.Sudoku {
	var vailSudoku *sudoku.Sudoku
	var resultSudoku *sudoku.Sudoku
	puzzle := basicSudoku.Solution()
	holeCounter := 0
	candidateHoles := randCandidateHoles()
	for _, hoIndex := range candidateHoles {
		holeCounter++
		old := puzzle[hoIndex]
		puzzle[hoIndex] = EMPTY
		vailSudoku = sudokuVerifyWithDfs(&puzzle)
		if vailSudoku == nil {
			puzzle[hoIndex] = old
			holeCounter--
		} else {
			resultSudoku = vailSudoku
		}

		if holeCounter >= digHoleTotal && resultSudoku != nil {
			return resultSudoku
		}
	}

	return nil
}

// this function use dfs algorithm verify with strict mode
// diffen way is sudokuVerifyWithDlx , that use [dlx] to solve puzzle,
// seem [dlx] way is faster solve very hard puzzle , but can't verify is one-solution puzzle
// if use [dlx] way to solve puzzle and use [dfs] to verify that will take more time
// so I remove sudokuVerifyWithDlx function , function just like :
// return sudoku.DLXSolve(*puzzle) == sudoku.SudokuGo2str(&solution)
func sudokuVerifyWithDfs(puzzle *[81]int8) *sudoku.Sudoku {
	vailSudoku := sudoku.Sudoku{}
	if err := vailSudoku.StrictInit(*puzzle); err != nil {
		return nil
	}
	return &vailSudoku
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func randCandidateHoles() []int {
	// 随机出 1-9 区的固定位置
	// 剩余 81 减 9 洗牌 作为候选 hole
	arr := make([]int, 81)
	for i := range arr {
		arr[i] = i
	}
	// make sure each zone must have one cell to fixed
	// need calculate random index on each zone , and sort them
	rand.Seed(time.Now().UnixNano())
	fixedPositionByZones := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	for i, zone := range fixedPositionByZones {
		x := rand.Intn(9)
		_, _, index := sudoku.LocationAtZone(zone, x)
		fixedPositionByZones[i] = index
	}
	for i, fixedPosition := range fixedPositionByZones {
		arr = remove(arr, fixedPosition-i)
	}
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	return arr
}
