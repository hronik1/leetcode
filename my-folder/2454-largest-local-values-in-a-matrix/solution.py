class Solution:
    def largestLocal(self, grid: List[List[int]]) -> List[List[int]]:
        length = len(grid)-2
        largest_locals = [[0]*length for i in range(length)]
        for i in range(1, len(grid)-1):
            for j in range(1, len(grid)-1):
                largest_locals[i-1][j-1] = self.largest_local_helper(grid, i, j)

        return largest_locals
    
    def largest_local_helper(self, grid, i, j):
        return max([grid[row][col] for row in range(i-1, i+2) for col in range(j-1, j+2)])

