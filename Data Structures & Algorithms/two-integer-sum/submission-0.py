class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        a = dict()
        for i in range(len(nums)):
            n = nums[i]
            t = target - n
            if t in a:
                return [a.get(t), i]
            a[n] = i
        