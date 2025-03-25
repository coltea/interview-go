class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        res = 0
        tmp = ""
        for i in s:
            if i not in tmp:
                tmp += i
            else:
                tmp = tmp[tmp.index(i) + 1:] + i
                # print(tmp)
            if len(tmp) > res:
                res = len(tmp)
        return res

