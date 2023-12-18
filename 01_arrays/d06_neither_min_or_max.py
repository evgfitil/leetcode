# https://leetcode.com/problems/neither-minimum-nor-maximum/

def solution(nums):
    if len(nums) <= 2:
        return -1
    fmin = fmax = nums[0]
    prev = -1
    for n in nums:
        if n < fmax and n > fmin:
            return n
        if n > fmax:
            prev = fmax
            fmax = n
        elif n < fmin:
            prev = fmin
            fmin = n
    return prev

nums = list(map(int, input().split()))
print(solution(nums))
