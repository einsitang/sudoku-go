package core

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type Sudoku struct {
	puzzle    [81]int8 `desc:"sudoku puzzle"`
	answer    [81]int8 `desc:"complete sudoku answer data"`
	rows      [9][9]bool
	cells     [9][9]bool
	zones     [9][9]bool
	beginTime time.Time
	endTime   time.Time
	nums      [9]int
}

func (_sudoku Sudoku) Puzzle() [81]int8 {
	return _sudoku.puzzle
}

func (_sudoku Sudoku) Answer() [81]int8 {
	return _sudoku.answer
}

func (_sudoku *Sudoku) Init(puzzle [81]int8) error {

	// initialize
	_sudoku.beginTime = time.Now()
	_sudoku.puzzle = puzzle
	_sudoku.answer = puzzle
	_sudoku.nums = ShuffleNumbers()
	rows, cells, zones := &_sudoku.rows, &_sudoku.cells, &_sudoku.zones

	defer func() {
		_sudoku.endTime = time.Now()
	}()

	for i, word := range puzzle {
		if word == -1 {
			continue
		}
		x, y, zone := Bearing(i)
		rows[x][word-1] = true
		cells[y][word-1] = true
		zones[zone][word-1] = true

	}

	// calculate
	isDone := _sudoku.calculate()
	if !isDone {
		return errors.New("puzzle can not be resolve")
	}
	return nil
}

func (_sudoku *Sudoku) Debug() {
	log.Println("--- debug sudoku info ---")
	log.Print("PUZZLE : \n", _sudoku.puzzleFormat())
	log.Print("ANSWER : \n", _sudoku.answerFormat())
	log.Printf("solved the puzzle with total time : %d ms", _sudoku.endTime.Sub(_sudoku.beginTime).Milliseconds())
}

func internalCalculate(_sudoku *Sudoku, index int) bool {

	if index >= 81 {
		return true
	}

	if _sudoku.answer[index] != -1 {
		return internalCalculate(_sudoku, index+1)
	}

	x, y, zone := Bearing(index)
	// log.Println(index, x, y, zone)
	rows, cells, zones := &_sudoku.rows, &_sudoku.cells, &_sudoku.zones
	for _, num := range _sudoku.nums {

		if !rows[x][num] && !cells[y][num] && !zones[zone][num] {
			rows[x][num] = true
			cells[y][num] = true
			zones[zone][num] = true
			_sudoku.answer[index] = int8(num + 1)
			if !internalCalculate(_sudoku, index+1) {
				_sudoku.answer[index] = -1
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

func (_sudoku *Sudoku) calculate() bool {

	_answer := &_sudoku.answer

	for i, word := range _answer {

		if word == -1 {
			return internalCalculate(_sudoku, i)
		}

	}

	return false
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
