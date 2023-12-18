# https://leetcode.com/problems/permutation-in-string/description/

def solution(s1, s2):
    s1map = dict()
    for i in s1:
        s1map[i] = s1map.get(i, 0) + 1
    s2map = dict()
    left, window_size = 0, len(s1)
    for right in range(len(s2)):
        current_char = s2[right]
        if current_char in s1map:
            s2map[current_char] = s2map.get(current_char, 0) + 1
        if right - left + 1 == window_size:
            if s1map == s2map:
                return True
            left_char = s2[left]
            if s2map.get(left_char):
                s2map[left_char] -= 1
                if s2map[left_char] == 0:
                    del s2map[left_char]
            left += 1
        right += 1

    return False


assert solution("ab", "eidbaooo") == True, "Test case 1 failed"
assert solution("ab", "eidboaoo") == False, "Test case 2 failed"
assert solution("abc", "cccccbabbbaaaa") == True, "Test case 3 failed"
assert solution("abc", "cccccbdbbbaaaa") == False, "Test case 4 failed"
assert solution("abc", "bbbca") == True, "Test case 5 failed"
assert solution("hello", "ooolleoooleh") == False, "Test case 6 failed"
assert solution("adc", "dcda") == True, "Test case 7 failed"
assert solution("r", "pilmtnzraxj") == True, "Test case 8 failed"
