# https://leetcode.com/problems/remove-element/

def solution(nums, val):
    l = 0
    for i in range(len(nums)):
        if nums[i] != val:
            nums[l] = nums[i]
            l += 1
    return l


nums = list(map(int, input().split()))
val = int(input())
print(solution(nums, val))
