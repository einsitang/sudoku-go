package main

import (
	"math/rand"
	"testing"
	"time"

	sudoku "github.com/einsitang/sudoku-go/core"
)

func TestLocationAtZone(t *testing.T) {
	zone, indexFromZone := 2, 5
	x, y, index := sudoku.LocationAtZone(zone, indexFromZone)
	t.Logf("zone[ %v ] , indexFromZone : %v : x : %v , y : %v , index : %v", zone, indexFromZone, x, y, index)
}

func TestIndexesAtZone(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	zone := rand.Intn(9)
	indexes := sudoku.IndexesAtZone(zone)
	t.Logf("zone[ %v ] : %v ", zone, indexes)
}
