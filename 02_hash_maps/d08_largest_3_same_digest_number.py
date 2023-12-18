# https://leetcode.com/problems/largest-3-same-digit-number-in-string/

def solution(num):
    prev_max = ""
    left, right = 0, 1
    while right < len(num) -1:
        while right < len(num) and num[left] == num[right] and right - left < 3:
            right += 1
        if right - left == 3:
            current_max = num[left:right]
            if current_max > prev_max:
                prev_max = current_max
        left = right
        right = left + 1
    return prev_max

num = str(input())
print(solution(num))
