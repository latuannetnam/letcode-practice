// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem:5. Longest Palindromic Substring
// Letcode link: https://leetcode.com/problems/longest-palindromic-substring/
// Level: medium
// Topics: String, Dynamic Programing
// Problem detail:
// Given a string s, return the longest palindromic substring in s.
// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.

package main

import "fmt"

// Runtime: 1133 ms, faster than 8.74% of Go online submissions for Longest Palindromic Substring.
// Memory Usage: 2.7 MB, less than 79.39% of Go online submissions for Longest Palindromic Substring.
func longestPalindromeBrustForce(s string) string {
	n := len(s)
	if n == 1 {
		return s
	}
	maxLen := 0
	longestS := ""
	for i := 0; i < n-1; i++ {
		fmt.Printf("Check longestPalindrome at index:%d\n", i)
		for j := i + 1; j <= n; j++ {
			isPalind := isPalindromeString(s[i:j])
			fmt.Printf(" index:%d subString:%s  isPalind:%t\n", i, s[i:j], isPalind)
			if isPalind && len(s[i:j]) > maxLen {
				maxLen = len(s[i:j])
				longestS = s[i:j]
			}
			fmt.Printf("  index:%d maxLen:%d longestS:%s\n", i, maxLen, longestS)
		}
	}
	return longestS
}

// Dynamic Programing: https://leetcode.com/problems/longest-palindromic-substring/solution/
// Runtime: 201 ms, faster than 25.70% of Go online submissions for Longest Palindromic Substring.
// Memory Usage: 8.4 MB, less than 17.76% of Go online submissions for Longest Palindromic Substring.
func longestPalindromeDP(s string) string {
	n := len(s)
	if n == 1 {
		return s
	}
	// DP State
	dp := make([][]bool, n)
	maxLen := 1
	longestS := string(s[0])
	for i := range dp {
		// Init state for dp
		dp[i] = make([]bool, n)
		dp[i][i] = true
		if i+1 < n {
			dp[i][i+1] = (s[i] == s[i+1])
			if dp[i][i+1] && maxLen < 2 {
				maxLen = 2
				longestS = s[i : i+2]
			}
		}

	}

	fmt.Printf("Begin check:  maxLen:%d longestS:%s\n", maxLen, longestS)
	fmt.Printf("%v\n", dp)

	// Build DP states
	// https://leetcode.com/problems/longest-palindromic-substring/discuss/2323248/C%2B%2B-or-Dynamic-Programming
	for i := 2; i < n; i++ {
		fmt.Printf("Check longestPalindrome at dp[row][%d]\n", i)
		for j := 0; j < n-i; j++ {
			dp[j][j+i] = dp[j+1][j+i-1] && (s[j] == s[j+i])
			fmt.Printf("dp[%d][%d]:%t -> dp[%d][%d]:%t\n", j+1, j+i-1, dp[j+1][j+i-1], j, j+i, dp[j][j+i])
			if dp[j][j+i] && maxLen < i+1 {
				maxLen = i + 1
				longestS = s[j : j+i+1]
			}
		}
		fmt.Printf("  maxLen:%d longestS:%s\n", maxLen, longestS)
	}

	return longestS

}

func mainLongestPalindrome() {
	s := "babad"
	// s = "cbbd"
	s = "ffffggggg"
	// s = "fffff"
	// s = "ffgg"
	// s = "fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffgggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggg"
	// s = "bkuadixhrtfehvbjjzojfyhuqyckdeilnftnklylyssbjivhvnplrzwrgcnrciypvacbhdnglflipnlpzgivjfieunhzeaytshrintdwhbvbkhbsqbjxwhqrlneiwkmccdnfcvapmtqihzyyoiaoqgtxkpbqqdboaxmmsdjxvdrwbhdnemqmsikoksfvjjovrbgdtfgmhgryvvpunzrsluqzibsvyubyhqevpnfnszzriljpmoevpqacbvdcsgfzmnkhnshsvynxxncqyjxqazcttkvjnkuvykgdrquybvlpwzladpetocuphzkgfuutqcbnttwjmkkwbmbidcyauopcxmsarodcqabirbawtlgsytlflsiolxjghjmqrymadpzaetcchyvkaezeavjemubbquclhcjcmbwxphllhfnzfyewpyyiirgnohylfbtfddeohtifrqiiwpdtyqjyemqjlnpcwlsylxjuxtnmcrpdznbzschmnjxnldxpkbrikpfsfwhsarrfjueubvnztlwmognapvxwelyvueheqxtncpxhzwplaxqqrbmfmmqhohucxinxidxzhndvstideuwrisjgpwkgvsdxmlnfgqzzksflmjzckkyutrwptfvcawfbvqxlztstpoitdepexxiqtsdtjmssqbdinalsqkjjqkgilbfxajninuclquszwbmstcdbywfhhnierqsegafyfqzvmqockcowfqwbgfxvdxbqobditvowhtdeptljetgj"
	fmt.Printf("s:%s reverse:%s\n", s, reverseString(s))

	// s1 := "bacdecab"
	// fmt.Printf("s1:%s isPalindrome:%t\n", s1, isPalindromeString(s1))

	fmt.Printf("Longest Palindrome:%s\n", longestPalindromeDP(s))

}
