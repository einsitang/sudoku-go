package generator

import (
	"errors"
	"runtime"

	sudoku "github.com/forfuns/sudoku-go/core"
)

type puzzleRule struct {
	Fill uint8 `desc:"each zone fill num"`
	Zone uint8 `desc:"how many zone much the fill rule"`
}

const (
	LEVEL_EASY   = 0
	LEVEL_MEDIUM = 1
	LEVEL_HARD   = 2
	LEVEL_EXPERT = 3

	EMPTY = -1
	FILL  = -2
)

func Generate(level int) (_sudoku sudoku.Sudoku, err error) {

	var rule [5]puzzleRule
	switch level {
	case LEVEL_EASY:
		rule = [5]puzzleRule{
			{7, 1},
			{6, 1},
			{5, 3},
			{4, 2},
			{3, 2},
		}
	case LEVEL_MEDIUM:
		rule = [5]puzzleRule{
			{6, 1},
			{5, 3},
			{4, 2},
			{3, 2},
			{2, 2},
		}
	case LEVEL_HARD:
		rule = [5]puzzleRule{
			{5, 1},
			{4, 2},
			{3, 3},
			{2, 2},
			{1, 1},
		}
	case LEVEL_EXPERT:
		rule = [5]puzzleRule{
			{5, 1},
			{4, 1},
			{3, 3},
			{2, 3},
			{1, 1},
		}
	default:
		err = errors.New("unknow level")
		return
	}
	ch := generateChan(rule)
	_sudoku = <-ch
	return
}

func generateChan(rule [5]puzzleRule) <-chan sudoku.Sudoku {
	var holes, simplePuzzle [81]int8

	// empty simple puzzle
	for i := range simplePuzzle {
		simplePuzzle[i] = EMPTY
	}

	// dig hole
	randZones := sudoku.ShuffleNumbers()
	countLoop := 0
	for _, r := range rule {
		// loop rule to make puzzle hole
		for x := 0; x < int(r.Zone); x++ {
			borehole(&holes, randZones[countLoop], r)
			countLoop++
		}

	}

	// get center indexes
	centerIndexes := func() [9]int {

		// calculate center coordinates
		var indexes [9]int
		i := 0
		for y := range [3]int{0, 1, 2} {
			_ = y
			for x := range [3]int{0, 1, 2} {
				indexes[i] = y*9 + 3 + x
				i++
			}
		}
		return indexes
	}()

	// random center simple puzzle
	nums := sudoku.ShuffleNumbers()
	for i, index := range centerIndexes {
		simplePuzzle[index] = int8(nums[i] + 1)
	}

	// concurrent generate
	n := runtime.NumCPU()
	if n < 5 {
		n = 5
	}
	ch := make(chan sudoku.Sudoku)
	// signal channel to make sure other goroutine will not block
	signal := make(chan int)
	for i := 0; i < n; i++ {
		go generate(ch, signal, holes, simplePuzzle)
	}
	signal <- 1
	return ch
}

func generate(ch chan<- sudoku.Sudoku, signal chan int, holes, basicPuzzle [81]int8) {
	basicSudoku := new(sudoku.Sudoku)
	basicSudoku.Init(basicPuzzle)
	puzzle := basicSudoku.Answer()

	// apply hole template to make new puzzle
	for i := range holes {
		if holes[i] != FILL {
			puzzle[i] = EMPTY
		}
	}

	// validate sudoku is good puzzle , make use twice calculate is the same answer
	var vs1, vs2 sudoku.Sudoku
	vs1 = sudoku.Sudoku{}
	vs2 = sudoku.Sudoku{}
	vs1.Init(puzzle)
	vs2.Init(puzzle)
	if sudoEquals(&vs1, &vs2) {
		// output final sudoku with success puzzle
		_, signalIsOk := <-signal
		if signalIsOk {
			ch <- vs1
			close(signal)
			close(ch)
		}

	} else {
		generate(ch, signal, holes, basicPuzzle)
	}
}

func sudoEquals(sudoku1, sudoku2 *sudoku.Sudoku) bool {
	ans1 := sudoku1.Answer()
	ans2 := sudoku2.Answer()
	for i := range ans1 {
		if ans1[i] != ans2[i] {
			return false
		}
	}

	return true
}

func borehole(puzzle *[81]int8, zoneIndex int, rule puzzleRule) {
	// hole for fill number
	fill := rule.Fill
	shuffleNums := sudoku.ShuffleNumbers()
	randIndees := shuffleNums[:fill]
	// make hold mark -> FILL = -2
	for _, randIndex := range randIndees {
		_, _, index := sudoku.BearingFromZone(zoneIndex, randIndex)
		puzzle[index] = FILL
	}
}
