# https://leetcode.com/problems/binary-subarrays-with-sum/description/

def solution(nums, goal):
    res = 0
    subarrays = list()
    left, right = 0, len(nums) -1
    while left <= right:
        if nums[left] + nums[right] == goal:
            res += 1
        right -= 1
    left, right = 0, len(nums) -1
    while left <= right:
        if nums[left] + nums[right] == goal:
            res += 1
        left += 1
    return res


nums = [0, 0, 0, 0, 0]
goal = 0
print(solution(nums, goal))
