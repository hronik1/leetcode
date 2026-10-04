class Solution:
    def maxArea(self, height: list[int]) -> int:
        if len(height) < 2:
            return 0
        
        start_i, end_i = 0, len(height) - 1
        best_volume = 0
        while start_i < end_i:
            start_height, end_height = height[start_i], height[end_i]
            cur_volume = (end_i - start_i) * min(start_height, end_height)
            if best_volume < cur_volume:
                best_volume = cur_volume
            
            if start_height < end_height:
                start_i += 1
            else:
                end_i -= 1

        return best_volume
