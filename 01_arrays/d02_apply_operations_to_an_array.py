# https://leetcode.com/problems/apply-operations-to-an-array/

def solution(nums):
    for i in range(len(nums)-1):
        if nums[i] == nums[i+1]:
            nums[i] *= 2
            nums[i+1] = 0
    l = 0
    for i in range(len(nums)):
        if nums[i] != 0:
            nums[l], nums[i] = nums[i], nums[l]
            l += 1
    return nums


nums = list(map(int, input().split()))
print(solution(nums))
