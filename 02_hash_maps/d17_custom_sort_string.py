# https://leetcode.com/problems/custom-sort-string/description/

def solution(order, s):
    smap = dict()
    res = []
    temp2 = []
    for n in s:
        smap[n] = smap.get(n, 0) + 1
    for char in order:
        if char in smap:
            temp = [char] * smap[char]
            res.extend(temp)
            del smap[char]
    for key in smap.keys():
        temp = key * smap[key]
        res.extend(temp)
    return "".join(res)


order = "hucw"
s = "utzoampdgkalexslxoqfkdjoczajxtuhqyxvlfatmptqdsochtdzgypsfkgqwbgqbcamdqnqztaqhqanirikahtmalzqjjxtqfnh"
print(solution(order, s))

assert solution("cba", "abcd") == "cbad", "Test 1 is failed"
assert solution("kqep", "pekeq") == "kqeep", "Test 2 is failed"
