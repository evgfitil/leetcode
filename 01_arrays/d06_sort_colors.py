# https://leetcode.com/problems/sort-colors/

def solution(nums):
    left, right, i = 0, len(nums) -1, 0
    while i <= right:
        if nums[i] == 2:
            nums[right], nums[i] = nums[i], nums[right]
            right -= 1
            if nums[i] == 1:
                i += 1
        elif nums[i] == 0:
            nums[left], nums[i] = nums[i], nums[left]
            i += 1
            left += 1
        else:
            i += 1
    return nums


nums = list(map(int, input().split()))
print(solution(nums))
