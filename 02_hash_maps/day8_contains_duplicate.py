# https://leetcode.com/problems/contains-duplicate/

def solution(nums):
    temp = dict()
    for i in nums:
        if temp.get(i, 0) >= 1:
            return True
        temp[i] = temp.get(i, 0) + 1
    return False


nums = list(map(int, input().split()))
print(solution(nums))
