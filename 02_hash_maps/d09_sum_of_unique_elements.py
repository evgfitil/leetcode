# https://leetcode.com/problems/sum-of-unique-elements/

def solution(nums):
    unique = dict()
    res = 0
    for num in nums:
        unique[num] = unique.get(num, 0) + 1
    for num in nums:
        if unique.get(num) == 1:
            res += num
    return res


nums = list(map(int, input().split()))
print(solution(nums))
