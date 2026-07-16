class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        a = set()
        value = len(a)
        for i in nums:
            a.add(i)
            if value == len(a):
                return True
            else:
                value = value + 1
        return False