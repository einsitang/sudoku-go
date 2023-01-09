package main

import (
	"encoding/json"
	generator "github.com/einsitang/sudoku-go/generator"
	"runtime"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {

	t.Logf("use goroutine : %v \n", runtime.NumGoroutine())
	beginTime := time.Now()
	sudoku, err := generator.Generate(generator.LEVEL_EXPERT)
	if err != nil {
		t.Fatal(err)
	}
	endTime := time.Now()
	t.Log("generate done")
	t.Logf("generated total time : %v ms", endTime.Sub(beginTime).Milliseconds())
	sudoku.Debug()
	bytes, _ := json.Marshal(sudoku.Puzzle())
	puzzleStr := string(bytes)
	t.Log("this is puzzle can be copy to the clipboard : ")
	t.Logf("%v", puzzleStr)
	t.Logf("the end goroutine : %v", runtime.NumGoroutine())
}

func BenchmarkGenerateLevelExpert(b *testing.B) {
	b.Logf("BenchmarkGenerateLevelExpert N : %v", b.N)
	beginTime := time.Now()
	for i := 0; i < b.N; i++ {
		_, err := generator.Generate(generator.LEVEL_EXPERT)
		if err != nil {
			b.Fatal(err)
		}
	}
	endTime := time.Now()
	b.Logf("generated (N:%v) total time : %v ms", b.N, endTime.Sub(beginTime).Milliseconds())
}

func BenchmarkGenerateLevelEasy(b *testing.B) {
	b.Logf("BenchmarkGenerateLevelEasy N : %v", b.N)
	beginTime := time.Now()
	for i := 0; i < b.N; i++ {
		_, err := generator.Generate(generator.LEVEL_EASY)
		if err != nil {
			b.Fatal(err)
		}
	}
	endTime := time.Now()
	b.Logf("generated (N:%v) total time : %v ms", b.N, endTime.Sub(beginTime).Milliseconds())
}
