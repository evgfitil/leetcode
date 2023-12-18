class NumArray:

    def __init__(self, nums: list[int]):
        self.nums = nums
        self.diff = [0] * (len(self.nums) + 1)
        self.diff[0] = self.nums[0]
        for i in range(1, len(nums)):
            self.diff[i] = self.nums[i] - self.nums[i - 1]

    def update(self, index: int, val: int) -> None:
        if index == 0:
            self.diff[0] == val
        else:
            self.diff[index] = val - self.nums[index - 1]
        
        self.nums[index] = val
        
        if index + 1 < len(self.nums):
            self.diff[index + 1] -= val - self.nums[index]


    def sumRange(self, left: int, right: int) -> int:
        self.left = left
        self.right = right
        
        res = 0
        while left <= right:
            res += self.nums[left]
            left += 1
        return res


# Your NumArray object will be instantiated and called as such:
# obj = NumArray(nums)
# obj.update(index,val)
# param_2 = obj.sumRange(left,right)
numArray = NumArray([1, 3, 5])
print(numArray.sumRange(0, 2))
print(numArray.update(1, 2))
print(numArray.sumRange(0, 2))
