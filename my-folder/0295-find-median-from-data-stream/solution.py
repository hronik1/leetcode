import heapq

class MedianFinder:

    def __init__(self):
        self.bottom_half = []
        self.top_half = []

    def addNum(self, num: int) -> None:
        if len(self.bottom_half) == len(self.top_half):
            smallest_top_half = heapq.heappushpop(self.top_half, num)
            heapq.heappush_max(self.bottom_half, smallest_top_half)
        else:
            largest_bottom_half = heapq.heappushpop_max(self.bottom_half, num)
            heapq.heappush(self.top_half, largest_bottom_half)
        
    def findMedian(self) -> float:
        if len(self.bottom_half) == len(self.top_half):
            return (self.bottom_half[0] + self.top_half[0]) / 2
        
        return self.bottom_half[0]


# Your MedianFinder object will be instantiated and called as such:
# obj = MedianFinder()
# obj.addNum(num)
# param_2 = obj.findMedian()
