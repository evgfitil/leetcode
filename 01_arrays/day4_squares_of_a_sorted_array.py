def solution(nums):
    res = [0] * len(nums)
    i, j = 0, len(nums) - 1
    n = len(nums) -1
    while i <= j:
        if abs(nums[i]) < abs(nums[j]):
            res[n] = nums[j] ** 2
            j -= 1
            n -= 1
        else:
            res[n] = nums[i] ** 2
            i += 1
            n -= 1
    return res


nums = list(map(int, input().split()))
print(solution(nums))
