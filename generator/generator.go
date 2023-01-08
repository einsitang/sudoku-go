package generator

import (
	"errors"
	"runtime"

	sudoku "github.com/einsitang/sudoku-go/core"
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

// Generate
// use this function to generate a sudoku quick , but not 100% signe answer sudoku ,
// if lower level LEVEL_MEDIUM will higher success rate
// whatever , you want quick and without care multi answer sudoku problem , use this is ok
func Generate(level int) (_sudoku sudoku.Sudoku, err error) {
	return doGenerate(level, false)
}

// StrictGenerate
// this function will make sure generate sudoku puzzle will not have multi answer , is 100% right sudoku,
// but it will take long time compared to Generate function
func StrictGenerate(level int) (_sudoku sudoku.Sudoku, err error) {
	return doGenerate(level, true)
}

func doGenerate(level int, strict bool) (_sudoku sudoku.Sudoku, err error) {

	enhance := (uint8)(0)
	rule := [5]puzzleRule{
		{7, 1},
		{6, 1},
		{5, 3},
		{4, 2},
		{3, 2},
	}
	switch level {
	case LEVEL_EASY:
		enhance = 0
	case LEVEL_MEDIUM:
		enhance = 8
	case LEVEL_HARD:
		enhance = 14
	case LEVEL_EXPERT:
		enhance = 18
	default:
		err = errors.New("unknown level , make sure range by [0,3]")
		return
	}
	ch := generateChan(rule, enhance, strict)
	_sudoku = <-ch
	return
}

func generateChan(rule [5]puzzleRule, enhance uint8, strict bool) <-chan sudoku.Sudoku {
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
	if n < 4 {
		n = 4
	}
	ch := make(chan sudoku.Sudoku)
	// signal channel to make sure other goroutine will not block
	signal := make(chan int)
	for i := 0; i < n; i++ {
		go generate(ch, signal, holes, enhance, strict, simplePuzzle)
	}
	signal <- 1
	return ch
}

//	the function to make cycle to validate sudoku calculate is same answer ,
//	and try many times will improve reliability
func repeatValidateSudoku(basicSudoku *sudoku.Sudoku, puzzle *[81]int8, tryCount int) (bool, *sudoku.Sudoku) {
	var vailSudoku sudoku.Sudoku
	isPass := true
	if tryCount < 0 {
		tryCount = 1
	}
	for tryCount > 0 {
		vailSudoku = sudoku.Sudoku{}
		_ = vailSudoku.Init(*puzzle)
		if !sudokuEquals(basicSudoku, &vailSudoku) {
			isPass = false
		}
		tryCount--
	}
	return isPass, &vailSudoku
}

func generate(ch chan<- sudoku.Sudoku, signal chan int, holes [81]int8, enhance uint8, strict bool, basicPuzzle [81]int8) {
	var vailSudoku sudoku.Sudoku
	basicSudoku := sudoku.Sudoku{}
	_ = basicSudoku.Init(basicPuzzle)
	puzzle := basicSudoku.Answer()

	// try dig hole and validate sudoku
	// apply hole template to make new puzzle
	basicDigHoleCounter := 0
	for i := range holes {
		if holes[i] != FILL {
			basicDigHoleCounter++
			puzzle[i] = EMPTY

			// basic dig hole less than mean puzzle resolve may not collision
			if basicDigHoleCounter < 5 {
				continue
			}

			isPass, matchSudoku := repeatValidateSudoku(&basicSudoku, &puzzle, basicDigHoleCounter/3)
			if !isPass {
				go generate(ch, signal, holes, enhance, strict, basicPuzzle)
				return
			}
			vailSudoku = *matchSudoku
		}
	}

	// deep dig hole and validate sudoku
	// 深度挖洞
	candidateHoles := buildCandidateHoles(holes)
	tryMoreHole := 1 + enhance
	for _, ho := range candidateHoles {
		for _, hoIndex := range ho {
			if tryMoreHole > 0 {

				old := puzzle[hoIndex]
				puzzle[hoIndex] = EMPTY

				if !strict {
					tryMoreHole--
					continue
				}

				// sudoku solver to make sure deep dig hole will match same answer
				isPass, matchSudoku := repeatValidateSudoku(&basicSudoku, &puzzle, 3)
				if isPass {
					vailSudoku = *matchSudoku
					tryMoreHole--
				} else {
					puzzle[hoIndex] = old
				}

			} else {
				break
			}
		}
	}

	if tryMoreHole > 0 {
		go generate(ch, signal, holes, enhance, strict, basicPuzzle)
	} else {
		_, signalIsOk := <-signal
		if signalIsOk {
			// if strict will not init valiSudoku to test , so need init by now
			//if reflect.ValueOf(vailSudoku).Kind() != reflect.Ptr {
			//	vailSudoku = sudoku.Sudoku{}
			//	_ = vailSudoku.Init(puzzle)
			//}
			//if !strict {
			//	vailSudoku = sudoku.Sudoku{}
			//	_ = vailSudoku.Init(puzzle)
			//}
			ch <- vailSudoku
			close(signal)
			close(ch)
		}
	}

	// validate sudoku is good puzzle , make sure twice calculate is the same answer
	//var vs1 sudoku.Sudoku
	//vs1 = sudoku.Sudoku{}
	//_ = vs1.Init(puzzle)
	//
	//_, signalIsOk := <-signal
	//if signalIsOk {
	//	log.Printf("尝试挖洞次数 : %v", digHoleFailCount)
	//	ch <- vs1
	//	close(signal)
	//	close(ch)
	//}
	//
	//tryCount := 8
	//done := true
	//
	//for i := 0; i < tryCount; i++ {
	//	vs2 = sudoku.Sudoku{}
	//	_ = vs2.Init(puzzle)
	//	if !sudoEquals(&vs1, &vs2) {
	//		done = false
	//		break
	//	}
	//	vs1 = vs2
	//}
	//
	//if done {
	//	_, signalIsOk := <-signal
	//	if signalIsOk {
	//		ch <- vs1
	//		close(signal)
	//		close(ch)
	//	}
	//} else {
	//	generate(ch, signal, holes, basicPuzzle)
	//}
}

func sudokuEquals(sudoku1, sudoku2 *sudoku.Sudoku) bool {
	ans1 := sudoku1.Answer()
	ans2 := sudoku2.Answer()
	for i := range ans1 {
		if ans1[i] != ans2[i] {
			//log.Printf("false : %v %v - %v", i, ans1[i], ans2[i])
			return false
		}
	}

	return true
}

func borehole(puzzle *[81]int8, zoneIndex int, rule puzzleRule) {
	// hole for fill number
	fill := rule.Fill
	shuffleNums := sudoku.ShuffleNumbers()
	randNums := shuffleNums[:fill]
	// make hold mark -> FILL = -2
	for _, randIndex := range randNums {
		_, _, index := sudoku.BearingFromZone(zoneIndex, randIndex)
		puzzle[index] = FILL
	}
}

func buildCandidateHoles(holes [81]int8) [9][]int {

	var candidateHoles [9][]int
	randZones := sudoku.ShuffleNumbers()
	for _, randZone := range randZones {
		// less keep 2 FILL position not to dig hole
		min := 2
		indexes := sudoku.IndexesFromZone(randZone)
		indexes = sudoku.Shuffle(indexes)

		arr := make([]int, 0)
		for _, index := range indexes {
			if holes[index] == FILL {
				if min <= 0 {

					arr = append(arr, index)
				}
				min--
			}
		}
		candidateHoles[randZone] = arr

	}
	return candidateHoles
}
