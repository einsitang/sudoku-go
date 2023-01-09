package core

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type Sudoku struct {
	puzzle            [81]int8 `desc:"sudoku puzzle"`
	answer            [81]int8 `desc:"complete sudoku answer data"`
	rows              [9][9]bool
	cells             [9][9]bool
	zones             [9][9]bool
	beginTime         time.Time
	endTime           time.Time
	nums              [9]int
	finishes          int
	isOneSolutionMode bool
}

func (_sudoku Sudoku) Puzzle() [81]int8 {
	return _sudoku.puzzle
}

func (_sudoku Sudoku) Answer() [81]int8 {
	return _sudoku.answer
}

func (_sudoku *Sudoku) StrictInit(puzzle [81]int8) error {
	_sudoku.isOneSolutionMode = true
	return _sudoku.Init(puzzle)
}

func (_sudoku *Sudoku) Init(puzzle [81]int8) error {

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

func (_sudoku *Sudoku) calculate() (err error) {

	_answer := &_sudoku.answer
	firstCheckPoint := 0
	for i, word := range _answer {
		if word == -1 {
			firstCheckPoint = i
			break
		}
	}

	if _sudoku.isOneSolutionMode {
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

func (_sudoku *Sudoku) Debug() {
	log.Println("--- debug sudoku info ---")
	log.Print("PUZZLE : \n", _sudoku.puzzleFormat())
	log.Print("ANSWER : \n", _sudoku.answerFormat())
	log.Printf("solved the puzzle with total time : %d ms", _sudoku.endTime.Sub(_sudoku.beginTime).Milliseconds())
}

func backtrackCalculate(_sudoku *Sudoku, index int) bool {

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

func (_self *Sudoku) dsfOneSolutionCalculate(_sudoku Sudoku, index int) {

	if _self.finishes > 1 {
		return
	}

	if index >= 81 {
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

			if index == 80 {
				_self.answer = _sudoku.answer
				_self.rows = _sudoku.rows
				_self.cells = _sudoku.cells
				_self.zones = _sudoku.zones
				_self.finishes++
				return
			}
			_self.dsfOneSolutionCalculate(_sudoku, index+1)
			answer[index] = -1
			rows[x][num] = false
			cells[y][num] = false
			zones[zone][num] = false
		}
	}
}

func (_sudoku *Sudoku) puzzleFormat() string {
	_puzzle := _sudoku.puzzle
	return sudokuFormat(_puzzle)
}

func (_sudoku *Sudoku) answerFormat() string {
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
		str = fmt.Sprintf(divide, str, word)
	}
	return str
}
