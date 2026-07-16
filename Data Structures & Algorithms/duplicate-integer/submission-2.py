class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        a = set()
        for i in nums:
            value = len(a)
            a.add(i)
            if value == len(a):
                return True
        return False