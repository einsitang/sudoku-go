package main

import (
	"log"
	"math/rand"
	"testing"
	"time"

	sudoku "github.com/einsitang/sudoku-go/core"
)

func TestShuffle(t *testing.T) {
	// nums := make([]int, 9)
	nums := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	for i := range nums {
		nums[i] = i
	}

	nums = sudoku.Shuffle(nums)
	log.Println(nums)
}

func TestBearingFromZone(t *testing.T) {
	zone, indexFromZone := 2, 5
	x, y, index := sudoku.BearingFromZone(zone, indexFromZone)
	log.Printf("zone[ %v ] , indexFromZone : %v : x : %v , y : %v , index : %v", zone, indexFromZone, x, y, index)
}

func TestIndexesFromZone(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	zone := rand.Intn(9)
	indexes := sudoku.IndexesFromZone(zone)
	log.Printf("zone[ %v ] : %v ", zone, indexes)
}
