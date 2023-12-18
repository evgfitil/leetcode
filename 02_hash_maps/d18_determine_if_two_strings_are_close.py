# https://leetcode.com/problems/determine-if-two-strings-are-close/description/

def solution(word1, word2):
    if len(word1) != len(word2):
        return False
    word1_map, word2_map = dict(), dict()
    for i in range(len(word1)):
        word1_map[word1[i]] = word1_map.get(word1[i], 0) + 1
        word2_map[word2[i]] = word2_map.get(word2[i], 0) + 1
    for i in range(len(word1)):
        if word1[i] not in word2_map or word2[i] not in word1_map:
            return False
    freq1, freq2 = dict(), dict()
    for f in word1_map.values():
        freq1[f] = freq1.get(f, 0) + 1
    for f in word2_map.values():
        freq2[f] = freq2.get(f, 0) + 1

    return freq1 == freq2


# word1, word2 = "abc", "bca"
# word1, word2 = "abc", "bcd"
# word1, word2 = "cabbba", "abbccc"
# word1, word2 = "cabbba", "aabbss"
word1, word2 = "abbzzca", "babzzcz"
print(solution(word1, word2))
