class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        d = dict()
        for sym in s:
            if sym in d:
                d[sym] += 1
            else:
                d[sym] = 1
        
        d2 = dict()
        for sym in t:
            if sym in d2:
                d2[sym] += 1
            else:
                d2[sym] = 1
        
        if d == d2:
            return True
        else:
            return False