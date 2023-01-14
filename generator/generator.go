package generator

import (
	"errors"
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
	// can't open because take too long time
	//LEVEL_HELL = 4

	// MIN_CONCURRENCY
	// minimum concurrency value is set here , if use concurrency to improve performance
	MIN_CONCURRENCY = 2

	// EMPTY puzzle empty value
	EMPTY = -1
)

// Generate
// this function will generate sudoku with one-solution
func Generate(level int) (_sudoku sudoku.Sudoku, err error) {
	// concurrent generate
	// if level below medium , just use one concurrent that will work fine
	n := runtime.NumCPU()
	n = n >> 1
	if n < MIN_CONCURRENCY {
		n = MIN_CONCURRENCY
	}

	digHoleTotal := 40
	switch level {
	case LEVEL_EASY:
		digHoleTotal = 40
		n = 1
	case LEVEL_MEDIUM:
		digHoleTotal = 45
		n = 1
	case LEVEL_HARD:
		digHoleTotal = 52
	case LEVEL_EXPERT:
		digHoleTotal = 56
	//case LEVEL_HELL:
	//	digHoleTotal = 60
	//	fmt.Printf("😈welcome to hell😈 this difficulty will take a long time...")
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
		go generate(sudokuCh, signal, digHoleTotal, &done)
	}
	signal <- 1
	_sudoku = <-sudokuCh
	return
}

func generate(sudokuCh chan<- sudoku.Sudoku, signal chan int, digHoleTotal int, done *bool) {
	if *done {
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

	var vailSudoku sudoku.Sudoku
	basicSudoku := sudoku.Sudoku{}
	_ = basicSudoku.Init(simplePuzzle)
	puzzle := basicSudoku.Answer()
	holeCounter := 0
	candidateHoles := randCandidateHoles()
	for _, hoIndex := range candidateHoles {
		old := puzzle[hoIndex]
		puzzle[hoIndex] = EMPTY
		vailSudoku = sudoku.Sudoku{}
		if err := vailSudoku.StrictInit(puzzle); err != nil {
			puzzle[hoIndex] = old
			continue
		}
		holeCounter++
		if holeCounter >= digHoleTotal {
			_, signalIsOpen := <-signal
			if signalIsOpen {
				*done = true
				sudokuCh <- vailSudoku
				close(signal)
				close(sudokuCh)
			}
			return
		}
	}
	if !*done {
		generate(sudokuCh, signal, digHoleTotal, done)
	}
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
		//fmt.Printf("fixedPosition : %v , i : %v len(arr) : %v \n", fixedPosition, i, len(arr))
		arr = remove(arr, fixedPosition-i)
	}
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	return arr
}
