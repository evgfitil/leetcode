# https://leetcode.com/problems/unique-number-of-occurrences/description/

def solution(arr):
    count = dict()
    for i in arr:
        count[i] = count.get(i, 0) + 1
    temp = set()
    for k, v in count.items():
        if v not in temp:
            temp.add(v)
        else:
            return False
    return True


arr = list(map(int, input().split()))
print(solution(arr))
