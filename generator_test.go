package main

import (
	"runtime"
	"testing"

	generator "github.com/forfuns/sudoku-go/generator"
)

func TestGenerate(t *testing.T) {

	t.Log(runtime.NumGoroutine())
	sudoku, err := generator.Generate(generator.LEVEL_EXPERT)
	if err != nil {
		t.Fatal(err)
	}
	sudoku.Debug()
	t.Log("generate done")
	t.Log(runtime.NumGoroutine())
}
