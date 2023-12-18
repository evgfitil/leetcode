# https://leetcode.com/problems/third-maximum-number/

def solution(nums):
    nums.sort(reverse=True)
    res = []
    count = 3
    left, right = 0, 0
    return nums



nums = list(map (int, input().split()))
print(solution(nums))
