# Letcode problem: https://leetcode.com/problems/longest-palindromic-substring/
# Level: medium
# Given a string s, return the longest palindromic substring in s
# Issue: timeout when run 150 test-cases
class Solution:
    def longestPalindrome2(self, s: str) -> str:
        length = 0
        longest = ""
        for idx in range(len(s)):
            for list_len in range(idx + length-1, len(s)):
                list_len += 1
                sub_list = s[idx:list_len]
                if len(sub_list) > length:
                    print("sub[{}:{}]".format(idx, list_len))
                    reversed_list = sub_list[::-1]
                    if sub_list == reversed_list:
                        length_temp = len(sub_list)
                        if length_temp > length:
                            length = length_temp
                            longest = sub_list
                        print("sub[{}:{}]:{} max:{}".format(
                            idx, list_len, sub_list, length))
        return longest

    def longestPalindrome3(self, s: str) -> str:
        length = 0
        longest = ""
        loops = 0
        if len(s) == 1:
            return s
        for idx in range(len(s)):
            list_len = len(s)
            while list_len > idx + length:
                loops += 1
                sub_list = s[idx:list_len]
                reversed_list = sub_list[::-1]
                if sub_list == reversed_list:
                    length_temp = len(sub_list)
                    if length_temp > length:
                        length = length_temp
                        longest = sub_list
                    print("loops:{} sub[{}:{}]:{} max:{}".format(
                        loops, idx, list_len, sub_list, length))
                # else:
                #     print("loops:{} sub[{}:{}]:{}".format(loops, idx, list_len, sub_list))
                list_len -= 1
        print("loops: {}  longest len:{}".format(loops, len(longest)))
        return longest

    def longestPalindrome(self, s: str) -> str:
        length = 0
        longest = ""
        loops = 0
        if len(s) == 1:
            return s
        idx = 0
        strlen = len(s)
        while idx < (strlen - length):
            list_len = len(s)
            while list_len > idx + length:
                loops += 1
                sub_list = s[idx:list_len]
                reversed_list = sub_list[::-1]
                if sub_list == reversed_list:
                    length_temp = len(sub_list)
                    if length_temp > length:
                        length = length_temp
                        longest = sub_list
                    print("loops:{} sub[{}:{}]:{} len:{} max:{}".format(
                        loops, idx, list_len,  sub_list, len(sub_list), length))
                # else:
                #     print("loops:{} sub[{}:{}]:{}  len:{} ".format(loops, idx, list_len,  sub_list, len(sub_list)))
                list_len -= 1
            idx += 1
        print("loops: {}  longest len:{}".format(loops, len(longest)))
        return longest


if __name__ == '__main__':
    solution = Solution()
    # s = "1212cdbabababcdefz"
    # s = "121221abcd"
    # s = "caaaaab"
    # s = "fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffgggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggg"
    s = "bkuadixhrtfehvbjjzojfyhuqyckdeilnftnklylyssbjivhvnplrzwrgcnrciypvacbhdnglflipnlpzgivjfieunhzeaytshrintdwhbvbkhbsqbjxwhqrlneiwkmccdnfcvapmtqihzyyoiaoqgtxkpbqqdboaxmmsdjxvdrwbhdnemqmsikoksfvjjovrbgdtfgmhgryvvpunzrsluqzibsvyubyhqevpnfnszzriljpmoevpqacbvdcsgfzmnkhnshsvynxxncqyjxqazcttkvjnkuvykgdrquybvlpwzladpetocuphzkgfuutqcbnttwjmkkwbmbidcyauopcxmsarodcqabirbawtlgsytlflsiolxjghjmqrymadpzaetcchyvkaezeavjemubbquclhcjcmbwxphllhfnzfyewpyyiirgnohylfbtfddeohtifrqiiwpdtyqjyemqjlnpcwlsylxjuxtnmcrpdznbzschmnjxnldxpkbrikpfsfwhsarrfjueubvnztlwmognapvxwelyvueheqxtncpxhzwplaxqqrbmfmmqhohucxinxidxzhndvstideuwrisjgpwkgvsdxmlnfgqzzksflmjzckkyutrwptfvcawfbvqxlztstpoitdepexxiqtsdtjmssqbdinalsqkjjqkgilbfxajninuclquszwbmstcdbywfhhnierqsegafyfqzvmqockcowfqwbgfxvdxbqobditvowhtdeptljetgj"
    print("String {}:{}".format(len(s), s))
    print("Longest:", solution.longestPalindrome(s))
    # print("Longest2:", solution.longestPalindrome3(s))
