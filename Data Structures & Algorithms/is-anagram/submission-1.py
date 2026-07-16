class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        d = dict()
        d2 = dict()
        for i in range(len(s)):
            if s[i] in d:
                d[s[i]] += 1
            else:
                d[s[i]] = 1
            
            if t[i] in d2:
                d2[t[i]] += 1
            else:
                d2[t[i]] = 1

        return d == d2