// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 238. Product of Array Except Self
// Letcode link: https://leetcode.com/problems/product-of-array-except-self/
// Level: medium
// Topics:Array, Prefix sum
// Problem detail:
//Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
// You must write an algorithm that runs in O(n) time and without using the division operation.

package main

import "fmt"

// Complexity: O(N^2)
// Time Limit Exceeded
func productExceptSelfBurstForce(nums []int) []int {
	product := []int{}
	for i := range nums {
		prefixProd := 1
		for j := range nums {
			if j != i {
				prefixProd *= nums[j]
			}
		}
		fmt.Printf("prefixProd[%d]:%d\n", i, prefixProd)
		product = append(product, prefixProd)
	}
	return product
}

// Complexity: O(N) with division
// Runtime: 33 ms, faster than 73.81% of Go online submissions for Product of Array Except Self.
// Memory Usage: 7.7 MB, less than 65.24% of Go online submissions for Product of Array Except Self.
func productExceptSelfBurstForce2(nums []int) []int {
	product := make([]int, len(nums))
	prefixProd := 1
	zeros := 0
	zeroIndex := -1
	for i := range nums {
		if nums[i] != 0 {
			prefixProd *= nums[i]
		} else {
			zeros++
			zeroIndex = i
		}
	}

	fmt.Printf("prefixProd:%d\n", prefixProd)

	if zeros > 1 {
		prefixProd = 0
	} else if zeros == 1 {
		product[zeroIndex] = prefixProd
	} else {
		for i := range nums {
			product[i] = prefixProd / nums[i]
		}
	}

	return product
}

// https://neetcode.io/
// Complexity: O(N) : Prefix and Suffix Product without division + space= O(N)
// Runtime: 28 ms, faster than 86.74% of Go online submissions for Product of Array Except Self.
// Memory Usage: 7.5 MB, less than 82.95% of Go online submissions for Product of Array Except Self.
func productExceptSelfPrefixSuffix(nums []int) []int {
	product := make([]int, len(nums))
	prefixProd := make([]int, len(nums)+2)
	suffixProd := make([]int, len(nums)+2)
	// Calculate prefix Product
	prefixProd[0] = 1
	for i := range nums {
		prefixProd[i+1] = prefixProd[i] * nums[i]
	}

	// Calculate subfix Product
	suffixProd[len(nums)+1] = 1
	for i := len(nums); i > 0; i-- {
		suffixProd[i] = suffixProd[i+1] * nums[i-1]
	}

	fmt.Printf("prefixProd:%v\n", prefixProd)
	fmt.Printf("suffixProd:%v\n", suffixProd)

	// Calculate prod[i] = prefix[i-1] * suffix[i+1]
	for i := range product {
		product[i] = prefixProd[i] * suffixProd[i+2]
	}
	return product
}

// https://neetcode.io/
// Complexity: O(N) : Prefix and Suffix Product without division + space= O(1)
func productExceptSelfPrefixSuffixSpace1(nums []int) []int {
	product := make([]int, len(nums))
	n := len(nums)

	// Calculate prefix Product
	prefix := 1
	for i := range nums {
		product[i] = prefix
		prefix *= nums[i]
	}
	fmt.Printf("prefixProd:%v\n", product)

	// Calculate postfix Product
	postfix := 1
	for i := n - 1; i >= 0; i-- {
		product[i] *= postfix
		postfix *= nums[i]
	}

	return product
}

func mainProductExceptSelf() {
	nums := []int{1, 2, 3, 4}
	// nums = []int{0, 0}
	// nums = []int{-1, 1, 0, -3, 3}
	// nums = []int{0, 4, 0}
	fmt.Printf("nums:%v\n", nums)
	// fmt.Printf("productExceptSelf:%v\n", productExceptSelfBurstForce2(nums))
	fmt.Printf("productExceptSelf:%v\n", productExceptSelfPrefixSuffixSpace1(nums))
}
