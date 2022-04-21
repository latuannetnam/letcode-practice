# Letcode problem:https://leetcode.com/problems/longest-substring-without-repeating-characters/
# Level: medium
# Given a string s, find the length of the longest substring without repeating characters.

class Solution:
    def max_non_repitive(self, sub_list):
        buff_list = {}
        if len(sub_list) == 1:
            return 1
        for item in sub_list:
            if item not in buff_list:
                buff_list[item] = item
            else:
                break
        return len(buff_list)

    def lengthOfLongestSubstring(self, s: str) -> int:
        lenght = 0
        str_list = [char for char in s]
        print("s:", str_list)
        for idx in range(len(str_list)):
            for list_lengh in range(idx+lenght-1, len(str_list)):
                list_lengh = list_lengh+1
                sub_list = str_list[idx:list_lengh]
                if len(sub_list) > lenght:
                    length_temp = self.max_non_repitive(sub_list)
                    print("sub_list[{}:{}]:{}:  max2:{}".format(
                        idx, list_lengh, sub_list, length_temp))
                    if lenght < length_temp:
                        lenght = length_temp
                    if length_temp < len(sub_list):
                        break
        return lenght

    def lengthOfLongestSubstring2(self, s: str) -> int:
        lenght = 0
        str_list = [char for char in s]
        print("s:", str_list)
        for idx in range(len(str_list)):
            sub_list = str_list[idx:]
            if len(sub_list) <= lenght:
                break
            length_temp = self.max_non_repitive(sub_list)
            print("sub_list[{}:]:{}:  max:{}".format(
                idx, sub_list, length_temp))
            if lenght < length_temp:
                lenght = length_temp
        return lenght


if __name__ == '__main__':
    solution = Solution()
    # s = "pwwkew"
    s = "abcabcbb"
    # s = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    print("String len:", len(s))
    print(solution.lengthOfLongestSubstring(s))
    # lst = [50, 70, 30, 20, 90, 10, 50]
    # print(lst[1:3])
