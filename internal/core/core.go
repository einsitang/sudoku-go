package core

import (
	"errors"
	"fmt"
	"log"
	"time"
)
type Sudoku interface {
	Puzzle() [81]int8
	Answer() [81]int8
	Solution() [81]int8
	Debug()
}

func Solve(puzzle [81]int8, option *SudokuOption) (Sudoku, error) {
	var err error
	_sudoku := &sudoku{option:option}
	if option.DLXMode {
		err = _sudoku.DLXInit(puzzle)
	} else {
		err = _sudoku.Init(puzzle)
	}
	return _sudoku, err
}

type SudokuOption struct {
	IsOneSolutionMode bool
	DLXMode           bool
}

type sudoku struct {
	puzzle            [81]int8 `desc:"sudoku puzzle"`
	answer            [81]int8 `desc:"complete sudoku answer data"`
	rows              [9][9]bool
	cells             [9][9]bool
	zones             [9][9]bool
	beginTime         time.Time
	endTime           time.Time
	nums              [9]int
	finishes          int
	option 			  *SudokuOption
}

func (_sudoku sudoku) Puzzle() [81]int8 {
	return _sudoku.puzzle
}

// Deprecated : seem answer is not explicit api , change to Solution
func (_sudoku sudoku) Answer() [81]int8 {
	return _sudoku.answer
}

func (_sudoku sudoku) Solution() [81]int8 {
	return _sudoku.answer
}

func (_sudoku *sudoku) StrictInit(puzzle [81]int8) error {
	_sudoku.option.IsOneSolutionMode = true
	return _sudoku.Init(puzzle)
}

func (_sudoku *sudoku) DLXInit(puzzle [81]int8) error {
	fmt.Println("use [dlx] caculate it : this is not ensure one-solution sudoku")
	_sudoku.beginTime = time.Now()
	_sudoku.puzzle = puzzle
	solutionStr := DLXSolve(puzzle)
	if len(solutionStr) != 81 {
		return errors.New("puzzle can not be resolve")
	}
	_sudoku.endTime = time.Now()
	_sudoku.finishes = 0
	_sudoku.answer = Str2sudokuGo(&solutionStr)

	return nil
}

func (_sudoku *sudoku) Init(puzzle [81]int8) error {

	// initialize
	_sudoku.beginTime = time.Now()
	_sudoku.puzzle = puzzle
	_sudoku.answer = puzzle
	_sudoku.nums = ShuffleNumbers()
	_sudoku.finishes = 0
	rows, cells, zones := &_sudoku.rows, &_sudoku.cells, &_sudoku.zones

	defer func() {
		_sudoku.endTime = time.Now()
	}()

	for i, word := range puzzle {
		if word == -1 {
			continue
		}
		x, y, zone := Location(i)
		rows[x][word-1] = true
		cells[y][word-1] = true
		zones[zone][word-1] = true

	}

	// calculate
	err := _sudoku.calculate()
	if err != nil {
		return errors.New("puzzle can not be resolve")
	}
	return nil
}

func (_sudoku *sudoku) calculate() (err error) {

	_answer := &_sudoku.answer
	firstCheckPoint := 0
	for i, word := range _answer {
		if word == -1 {
			firstCheckPoint = i
			break
		}
	}

	if _sudoku.option.IsOneSolutionMode {
		_sudoku.dsfOneSolutionCalculate(*_sudoku, firstCheckPoint)
		if _sudoku.finishes > 1 {
			err = errors.New("puzzle is not one-solution sudoku")
			return
		}
		for i := 0; i < 81; i++ {
			if _sudoku.answer[i] == -1 {
				//_sudoku.Debug()
				err = errors.New("puzzle can't solve")
				return
			}
		}
		return
	}

	if !backtrackCalculate(_sudoku, firstCheckPoint) {
		err = errors.New("puzzle can't solve")
	}
	return
}

func (_sudoku *sudoku) Debug() {
	log.Println("--- debug sudoku info ---")
	log.Print("PUZZLE : \n", _sudoku.puzzleFormat())
	log.Print("SOLUTION : \n", _sudoku.answerFormat())
	log.Printf("solved the puzzle with total time : %d ms", _sudoku.endTime.Sub(_sudoku.beginTime).Milliseconds())
}

func backtrackCalculate(_sudoku *sudoku, index int) bool {

	if index >= 81 {
		return true
	}

	if _sudoku.answer[index] != -1 {
		return backtrackCalculate(_sudoku, index+1)
	}

	x, y, zone := Location(index)
	rows, cells, zones := &_sudoku.rows, &_sudoku.cells, &_sudoku.zones
	answer := &_sudoku.answer
	for _, num := range _sudoku.nums {

		if !rows[x][num] && !cells[y][num] && !zones[zone][num] {
			rows[x][num] = true
			cells[y][num] = true
			zones[zone][num] = true
			answer[index] = int8(num + 1)
			if !backtrackCalculate(_sudoku, index+1) {
				answer[index] = -1
				rows[x][num] = false
				cells[y][num] = false
				zones[zone][num] = false
			} else {
				return true
			}
		}
	}

	return false

}

func (_self *sudoku) dsfOneSolutionCalculate(_sudoku sudoku, index int) {

	if _self.finishes > 1 {
		return
	}

	if index >= 81 {
		for _, num := range _sudoku.answer {
			if num == -1 {
				return
			}
		}
		_self.answer = _sudoku.answer
		_self.rows = _sudoku.rows
		_self.cells = _sudoku.cells
		_self.zones = _sudoku.zones
		_self.finishes++
		return
	}

	if _sudoku.answer[index] != -1 {
		_self.dsfOneSolutionCalculate(_sudoku, index+1)
		return
	}

	x, y, zone := Location(index)
	rows, cells, zones := &_sudoku.rows, &_sudoku.cells, &_sudoku.zones
	answer := &_sudoku.answer
	for _, num := range _sudoku.nums {

		if !rows[x][num] && !cells[y][num] && !zones[zone][num] {
			rows[x][num] = true
			cells[y][num] = true
			zones[zone][num] = true
			answer[index] = int8(num + 1)

			_self.dsfOneSolutionCalculate(_sudoku, index+1)
			answer[index] = -1
			rows[x][num] = false
			cells[y][num] = false
			zones[zone][num] = false
		}
	}
}

func (_sudoku *sudoku) puzzleFormat() string {
	_puzzle := _sudoku.puzzle
	return sudokuFormat(_puzzle)
}

func (_sudoku *sudoku) answerFormat() string {
	_answer := _sudoku.answer
	return sudokuFormat(_answer)
}

func sudokuFormat(puzzle [81]int8) string {
	var str, divide string
	str = ""
	for i, word := range puzzle {
		if i != 0 && i%9 == 0 {
			str = fmt.Sprintf("%s\n", str)
			if i/9%3 == 0 {
				str = fmt.Sprintf("%s\n", str)
			}
		}
		if i%3 == 0 {
			// two blank space
			divide = "%s  %2v"
		} else {
			// one blan space
			divide = "%s %2v"
		}
		var wordStr string
		if word == -1 {
			wordStr = "·"
		} else {
			wordStr = fmt.Sprint(word)
		}
		str = fmt.Sprintf(divide, str, wordStr)
	}
	return str
}
