# Author: Le Anh Tuan (latuannetnam@gmail.com)
# Letcode problem: https://leetcode.com/problems/longest-common-prefix/
# Level: medium
# Write a function to find the longest common prefix string amongst an array of strings.
# If there is no common prefix, return an empty string "".

# Input: strs = ["flower","flow","flight"]
# Output: "fl"

class Solution:
     def longestCommonPrefix(self, strs) -> str:
        minlen = len(strs[0])
        maxlen = 0
        for str in strs:
            minlen = min(minlen, len(str))
            maxlen = max(maxlen, len(str))
        print("min len:{} max len:{}".format(minlen, maxlen))
        
        prefix = ""
        prefix_len = minlen
        while prefix_len > 0:
            prefix_temp = strs[0][0:prefix_len]
            print("prefix_len:{} prefix temp:{}".format(prefix_len, prefix_temp))
            for str in strs:
                if str[0:prefix_len] != prefix_temp:
                    prefix_temp = ""
            if prefix_temp != "":
                prefix = prefix_temp
                break
            prefix_len -= 1    

        return prefix


if __name__ == '__main__':
    
    solution = Solution()
    # strs = ["flower","flow","flight"]
    strs = ["dog","racecar","car"]
    print("Input:{} prefix:{}".format(strs, solution.longestCommonPrefix(strs)))
    
