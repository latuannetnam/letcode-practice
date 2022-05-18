// Author: Le Anh Tuan (latuannetnam@gmail.com)
// 283. Move Zeroes
// Letcode problem:
// Level: easy
// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
// Note that you must do this in-place without making a copy of the array.
// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]

package main

import "fmt"

// Runtime: 165 ms, faster than 5.03% of Go online submissions for Move Zeroes.
// Memory Usage: 7.3 MB, less than 27.51% of Go online submissions for Move Zeroes.
func moveZeroesSlow(nums []int) {
	if len(nums) == 1 {
		return
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == 0 && nums[j] != 0 {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

// Runtime: 28 ms, faster than 53.16% of Go online submissions for Move Zeroes.
// Memory Usage: 6.8 MB, less than 56.25% of Go online submissions for Move Zeroes.
func moveZeroes(nums []int) {
	if len(nums) == 1 {
		return
	}
	z_pos := 0
	nz_pos := 0
	for z_pos < len(nums) && nz_pos < len(nums) {
		if nums[z_pos] != 0 {
			z_pos++
		}
		if nums[nz_pos] == 0 {
			nz_pos++
		}

		fmt.Printf("z_pos:%d nz_pos:%d\n", z_pos, nz_pos)
		if z_pos < len(nums) && nz_pos < len(nums) {
			fmt.Printf("nums[z_pos:%d]:%d nums[nz_pos:%d]:%d\n", z_pos, nums[z_pos], nz_pos, nums[nz_pos])
			if nums[z_pos] == 0 && nums[nz_pos] != 0 {
				if z_pos < nz_pos {
					nums[z_pos], nums[nz_pos] = nums[nz_pos], nums[z_pos]
					z_pos++
				}
				nz_pos++
			}
		}
	}

}

func mainMoveZeroes() {
	var nums []int
	nums = []int{0, 1, 0, 3, 12}
	fmt.Printf("Nums: %v\n", nums)
	moveZeroes(nums)
	fmt.Printf("Nums after: %v\n", nums)

	nums = []int{0, 1}
	fmt.Printf("Nums: %v\n", nums)
	moveZeroes(nums)
	fmt.Printf("Nums after: %v\n", nums)

	nums = []int{0, 0}
	fmt.Printf("Nums: %v\n", nums)
	moveZeroes(nums)
	fmt.Printf("Nums after: %v\n", nums)

	nums = []int{1, 0, 0}
	fmt.Printf("Nums: %v\n", nums)
	moveZeroes(nums)
	fmt.Printf("Nums after: %v\n", nums)

	nums = []int{1, 0, 1}
	fmt.Printf("Nums: %v\n", nums)
	moveZeroes(nums)
	fmt.Printf("Nums after: %v\n", nums)
}
