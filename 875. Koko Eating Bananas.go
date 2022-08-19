// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 875. Koko Eating Bananas
// Letcode link: https://leetcode.com/problems/koko-eating-bananas/
// Level: medium
// Topics: Array, Binary Search
// Problem detail:
// Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone and will come back in h hours.
// Koko can decide her bananas-per-hour eating speed of k.
// Each hour, she chooses some pile of bananas and eats k bananas from that pile.
// If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.
// Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.
// Return the minimum integer k such that she can eat all the bananas within h hours.

package main

import (
	"fmt"
	"sort"
)

// Runtime: 105 ms, faster than 6.69% of Go online submissions for Koko Eating Bananas.
// Memory Usage: 7.7 MB, less than 5.58% of Go online submissions for Koko Eating Bananas.
func minEatingSpeedBinarySearch1(piles []int, h int) int {
	sort.Ints(piles)
	n := len(piles)

	calcPileTime := func(pile int) int {
		time := 0
		// fmt.Printf(" calc pileTime for:%d\n", pile)
		for i := range piles {
			time += piles[i] / pile
			remain := piles[i] % pile
			if remain > 0 {
				time++
			}
			// fmt.Printf("  pile[%d]:%d->%d\n", i, piles[i], time)
		}
		return time
	}
	// Calculate total time to eat if each time Kobo eat exactly piles[index]
	var left, right int
	pileTime1 := calcPileTime(piles[0])
	pileTime2 := calcPileTime(piles[n-1])
	if pileTime1 == h {
		return piles[0]
	} else if pileTime2 == h {
		return piles[n-1]
	} else if pileTime1 > h && pileTime2 < h {
		left = piles[0]
		right = piles[n-1]
	} else if pileTime1 < h {
		left = 1
		right = piles[0]
	} else {
		return 0
	}

	minPile := right
	for left < right {
		mid := left + (right-left)/2
		pileTime := calcPileTime(mid)
		fmt.Printf("Left:%d Right:%d Mid[%d]:%d/%d\n", left, right, mid, pileTime, h)
		if pileTime <= h {
			if mid < minPile {
				minPile = mid
			}
			right = mid
		} else if pileTime > h {
			left = mid + 1
		}
	}

	return minPile
}

// Neetcode
// Runtime: 72 ms, faster than 42.75% of Go online submissions for Koko Eating Bananas.
// Memory Usage: 7.3 MB, less than 36.80% of Go online submissions for Koko Eating Bananas.
func minEatingSpeed(piles []int, h int) int {
	sort.Ints(piles)
	n := len(piles)

	calcPileTime := func(pile int) int {
		time := 0
		// fmt.Printf(" calc pileTime for:%d\n", pile)
		for i := range piles {
			time += (piles[i]-1)/pile + 1

			// fmt.Printf("  pile[%d]:%d->%d\n", i, piles[i], time)
		}
		return time
	}
	// Calculate total time to eat if each time Kobo eat exactly piles[index]
	var left, right int
	pileTime1 := calcPileTime(piles[0])
	pileTime2 := calcPileTime(piles[n-1])
	if pileTime1 == h {
		return piles[0]
	} else if pileTime2 == h {
		return piles[n-1]
	} else if pileTime1 > h && pileTime2 < h {
		left = piles[0]
		right = piles[n-1]
	} else if pileTime1 < h {
		left = 1
		right = piles[0]
	} else {
		return 0
	}

	minPile := right
	for left < right {
		mid := left + (right-left)/2
		pileTime := calcPileTime(mid)
		fmt.Printf("Left:%d Right:%d Mid[%d]:%d/%d\n", left, right, mid, pileTime, h)
		if pileTime <= h {
			minPile = mid
			right = mid
		} else {
			left = mid + 1
		}
	}

	return minPile
}

func mainMinEatingSpeed() {
	piles := []int{3, 6, 7, 11}
	h := 8

	piles = []int{30, 11, 23, 4, 20}
	h = 6
	// piles = []int{30}
	// h = 5

	// piles = []int{312884470}
	// h = 312884469
	fmt.Printf("Piles:%v hours:%d\n", piles, h)
	fmt.Printf("min k:%d\n", minEatingSpeed(piles, h))
}
