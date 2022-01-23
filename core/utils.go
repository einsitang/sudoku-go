package core

import (
	"math/rand"
	"time"
)

func Bearing(index int) (x, y, zone int) {
	x = index / 9
	y = index % 9
	zone = x/3*3 + y/3
	// log.Printf("bearing index %v ,x %v, y %v, zone %v", index, x, y, zone)
	return
}

func BearingFromZone(zone, indexFromZone int) (x, y, index int) {
	x = zone/3*3 + indexFromZone/3
	y = zone%3*3 + indexFromZone%3
	index = x*9 + y
	return
}

func Shuffle(nums [9]int) [9]int {
	rand.Seed(time.Now().UnixNano())
	for i := range nums {
		n := rand.Intn(9)
		nums[i], nums[n] = nums[n], nums[i]
	}
	return nums
}

func ShuffleNumbers() [9]int {
	shuffleNums := Shuffle([9]int{0, 1, 2, 3, 4, 5, 6, 7, 8})
	return shuffleNums
}
