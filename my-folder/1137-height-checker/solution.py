class Solution:
    def heightChecker(self, heights: List[int]) -> int:
        expected_heights = sorted(heights)
        wrong = 0
        for i in range(len(heights)):
            if expected_heights[i] != heights[i]:
                wrong += 1

        return wrong
        
