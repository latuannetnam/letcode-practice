# Author: Le Anh Tuan (latuannetnam@gmail.com)
# Letcode problem: https://leetcode.com/problems/roman-to-integer/
# Level: easy
# Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M
# Given a roman numeral, convert it to an integer.
# Symbol       Value
# I             1
# III           3
# IV            4
# V             5
# VIII          8
# IX            9
# X             10
# XIII          13
# XXXIX         39
# XL            40
# XLVIII        48
# L             50
# LXXXIX        89
# XC            90
# XCVIII        98
# C             100
# CCCLXXIII     373
# CD            400
# CDXXVIII      428
# D             500
# DCCCLXIX       869
# CM            900
# CMLCIX        999
# M             1000
# MMMCCCXXXIII  3333 

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
    def romanToInt(self, s: str) -> int:
        num = 0
        index = 0
        while index < len(s):
            #  1000
            if s[index] == 'M':
                num += 1000
            elif s[index] == 'C':
                if index + 1 < len(s):
                    # 900
                    if s[index+1] == 'M':
                        num += 900
                        index +=1
                    # 400
                    elif s[index+1] == 'D':
                        num += 400
                        index +=1
                    else:
                    # 100
                        num += 100
                else:
                    # 100
                        num += 100
            elif s[index] == 'D':
                # 500
                    num += 500
            elif s[index] == 'X':
                if index + 1 < len(s):
                    if s[index+1] == 'C':
                        num += 90
                        index +=1
                    elif s[index+1] == 'L':
                        num += 40
                        index +=1
                    else:
                        num += 10
                else:
                    num += 10
            elif s[index] == 'L':
                num += 50
            elif s[index] == 'V':
                num += 5    
            elif s[index] == 'I':
                if index + 1 < len(s):
                    if s[index+1] == 'X':
                        num += 9
                        index += 1
                    elif s[index+1] == 'V':
                        num += 4
                        index += 1
                    else:
                        num += 1
                else:
                    num += 1 

            print('index:{} num:{}'.format(index, num))
            index += 1
        return num

    # Best solution: https://leetcode.com/problems/roman-to-integer/discuss/1985151/Python3-Solution-95-runtime-78-space
    def romanToInt2(self, s: str) -> int:
        SYMBOLS = {'I' : 1,
                'V' : 5,
                'X' : 10,
                'L' : 50,
                'C' : 100,
                'D' : 500,
                'M' : 1000}
        
        solution = 0
        
        # s[sym + 1] will be out of array bounds if you loop to len(s)
        for sym in range(0, len(s) - 1):
            if SYMBOLS[ s[sym] ] < SYMBOLS[ s[sym + 1]]:
                solution -= SYMBOLS[ s[sym] ]
            else:
                solution += SYMBOLS[ s[sym] ]
        
        solution += SYMBOLS[ s[-1] ]     # account for last character
        
        return solution
        


if __name__ == '__main__':
    
    solution = Solution()
    # roman = "MCMXCIV"
    # roman = "IV"
    # roman = "LVIII"
    # roman = "IX"
    # roman = "MDLXX"
    roman = "MMMCC"
    print("roman:{}  num:{}".format(roman, solution.romanToInt(roman)))
    
