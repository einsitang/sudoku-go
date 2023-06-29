package core

import (
	"math/rand"
	"time"
)

const (
	CONST_EASY_HOLES   = 40
	CONST_MEDIUM_HOLES = 45
	CONST_HARD_HOLES   = 50
	CONST_EXPERT_HOLES = 56
	// ⚠️ hell is realy hard and very consumptive performance ⚠️
	CONST_HELL_HOLES = 60
)

func Location(index int) (x, y, zone int) {
	x = index / 9
	y = index % 9
	zone = x/3*3 + y/3
	return
}

func LocationAtZone(zone, indexFromZone int) (x, y, index int) {
	x = zone/3*3 + indexFromZone/3
	y = zone%3*3 + indexFromZone%3
	index = x*9 + y
	return
}

func IndexesAtZone(zone int) [9]int {
	var indexes [9]int
	i := 0
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			indexes[i] = ((y + (zone/3)*3) * 9) + (x + (zone%3)*3)
			i++
		}
	}
	return indexes
}

func Shuffle(nums [9]int) [9]int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	return nums
}

func ShuffleNumbers() [9]int {
	shuffleNums := Shuffle([9]int{0, 1, 2, 3, 4, 5, 6, 7, 8})
	return shuffleNums
}
