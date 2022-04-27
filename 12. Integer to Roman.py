# Author: Le Anh Tuan (latuannetnam@gmail.com)
# Letcode problem: https://leetcode.com/problems/integer-to-roman/
# Level: medium
# Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M
# Given an integer, convert it to a roman numeral.
# Symbol       Value
# I             1
# V             5
# X             10
# L             50
# C             100
# D             500
# M             1000

# Input: num = 58
# Output: "LVIII"
# Explanation: L = 50, V = 5, III = 3
# 4 = IV
# 12 = XII
# 27 = XXVII
# 1994 = MCMXCIV
# Roman numerals are usually written largest to smallest from left to right.
# However, the numeral for four is not IIII. Instead, the number four is written
# as IV. Because the one is before the five we subtract it making four. The same
# principle applies to the number nine, which is written as IX. There are six
# instances where subtraction is used:
# I can be placed before V (5) and X (10) to make 4 and 9.
# X can be placed before L (50) and C (100) to make 40 and 90.
# C can be placed before D (500) and M (1000) to make 400 and 900.

class Solution:
    def intToRoman(self, num: int) -> str:
        roman = ""
        level = 0
        while num > 0:
            digit = num % 10
            level += 1
            if level == 1:
                if digit <= 3:
                    roman = 'I' * digit
                elif digit == 4:
                    roman = 'IV'
                elif digit == 5:
                    roman = 'V'
                elif digit < 9:
                    roman = 'V' + (digit-5) * 'I'
                else:
                    roman = 'IX'
            elif level == 2:
                if digit <= 3:
                    roman = 'X' * digit + roman
                elif digit == 4:
                    roman = 'XL' + roman
                elif digit == 5:
                    roman = 'L' + roman
                elif digit < 9:
                    roman = 'L' + (digit-5) * 'X' + roman
                else:
                    roman = 'XC' + roman
            elif level == 3:
                if digit <= 3:
                    roman = 'C' * digit + roman
                elif digit == 4:
                    roman = 'CD' + roman
                elif digit == 5:
                    roman = 'D' + roman
                elif digit < 9:
                    roman = 'D' + (digit-5) * 'C' + roman
                else:
                    roman = 'CM' + roman
            elif level == 4:
                if digit <= 3:
                    roman = 'M' * digit + roman

            print("num:{} level:{} digit:{} roman:{}".format(
                num, level, digit, roman))
            num //= 10

        return roman


if __name__ == '__main__':
    solution = Solution()
    num = 1994
    print("number:{}  roman:{}".format(num, solution.intToRoman(num)))
