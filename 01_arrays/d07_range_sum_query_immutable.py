## simple
# class NumArray:

#     def __init__(self, nums: list[int]):
#         self.nums = nums
        

#     def sumRange(self, left: int, right: int) -> int:
#         self.left = left
#         self.right = right
#         res = 0
#         while left <= right:
#             res += self.nums[left]
#             left += 1
#         return res

class NumArray:

    def __init__(self, nums: list[int]):
        self.nums = nums
        self.prefix_sums = [sum(self.nums[:i+1]) for i in range(len(self.nums))]


    def sumRange(self, left: int, right: int) -> int:
        self.left = left
        self.right = right
        if self.left == 0:
            return self.prefix_sums[right]
        else:
            self.left -= 1
            return self.prefix_sums[self.right] - self.prefix_sums[self.left]

# Your NumArray object will be instantiated and called as such:
# obj = NumArray(nums)
# param_1 = obj.sumRange(left,right)

numArray = NumArray([-2, 0, 3, -5, 2, -1])
print(numArray.sumRange(0, 2))
print(numArray.sumRange(2, 5))
print(numArray.sumRange(0, 5))
