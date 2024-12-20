package sudoku

import (
	core "github.com/einsitang/sudoku-go/v2/internal/core"
	generator "github.com/einsitang/sudoku-go/v2/internal/generator"
)

const (
	LEVEL_EASY   = generator.LEVEL_EASY
	LEVEL_MEDIUM = generator.LEVEL_MEDIUM
	LEVEL_HARD   = generator.LEVEL_HARD
	LEVEL_EXPERT = generator.LEVEL_EXPERT
	// LEVEL_HELL
	// ⚠️ this difficulty will take a long time , carefully to use
	LEVEL_HELL = generator.LEVEL_HELL
)

type OptionFunc func(option *core.SudokuOption)

func WithStrict() OptionFunc {
	return func(option *core.SudokuOption) {
		option.IsOneSolutionMode = true
	}
}

func WithDLX() OptionFunc {
	return func(option *core.SudokuOption) {
		option.IsOneSolutionMode = false
		option.DLXMode = true
	}
}
func Solve(puzzle [81]int8, opts ...OptionFunc) (core.Sudoku, error) {
	option := &core.SudokuOption{}

	for _, opt := range opts {
		opt(option)
	}

	return core.Solve(puzzle,option)
}
func Generate(level int) (core.Sudoku, error) {
	return generator.Generate(level)
}
