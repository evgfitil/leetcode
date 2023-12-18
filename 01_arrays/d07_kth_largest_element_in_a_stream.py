import heapq

class KthLargest:

    def __init__(self, k: int, nums: list[int]):
        self.heap = nums
        self.k = k
        heapq.heapify(self.heap) # Transform list x into a heap, in-place, in linear time

        while len(self.heap) > k:
            heapq.heappop(self.heap)
        

    def add(self, val: int) -> int:
        heapq.heappush(self.heap, val)
        if len(self.heap) > self.k:
            heapq.heappop(self.heap)
        return self.heap[0]


# Your KthLargest object will be instantiated and called as such:
# obj = KthLargest(k, nums)
# param_1 = obj.add(val)

kthLargest = KthLargest(3, [4, 5, 8, 2])
print(kthLargest.add(3))  # Выведет 4
print(kthLargest.add(5))  # Выведет 5
print(kthLargest.add(10)) # Выведет 5
print(kthLargest.add(9))  # Выведет 8
print(kthLargest.add(4))  # Выведет 8
