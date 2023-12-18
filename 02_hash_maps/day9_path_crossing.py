# https://leetcode.com/problems/path-crossing/description/

def solution(path):
    start = (0, 0)
    paths = set()
    paths.add(start)
    for p in path:
        if p == "N":
            start = (start[0] + 1, start[1])
        elif p == "S":
            start = (start[0] - 1, start[1])
        elif p == "E":
            start = (start[0], start[1] + 1)
        elif p == "W":
            start = (start[0], start[1] - 1)
        
        if start in paths:
            return True
        else:
            paths.add(start)

    return False



path = str(input())
print(solution(path))
