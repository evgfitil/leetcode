# https://leetcode.com/problems/destination-city/description/

# set solution
def solution(paths):
    sources = set()
    for path in paths:
        sources.add(path[0])
    for path in paths:
        if path[1] not in sources:
            return path[1]
        

paths = [["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]]
print(solution(paths))
