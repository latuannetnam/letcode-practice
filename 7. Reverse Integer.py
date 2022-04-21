# Letcode problem: https://leetcode.com/problems/reverse-integer/
# Level: medium
# Given a signed 32-bit integer x, return x with its digits reversed.
# If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0

from re import S
import re


class Solution:
    def reverse(self, x: int) -> int:
        max = pow(2, 31) - 1
        min = pow(2, 31)*-1
        num = abs(x)
        reversed_num = 0
        while num != 0:
            digit = num % 10
            reversed_num = reversed_num * 10 + digit
            num //= 10

        if x < 0:
            reversed_num *= -1
        if reversed_num > max or reversed_num < min:
            reversed_num = 0
        return reversed_num


if __name__ == '__main__':
    x = 1534236469
    solution = Solution()
    print('x:{} reverse:{}'.format(x, solution.reverse(x)))
