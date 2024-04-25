class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        num_islands = 0
        seen = set()
        for row in range(len(grid)):
            for col in range(len(grid[row])):
                cell = (row, col)
                if cell not in seen and grid[row][col] == '1':
                    self.search(grid, seen, row, col)
                    num_islands += 1
                
        return num_islands
        
    def search(self, grid, seen, row, col):
        cell = (row, col)
        if cell in seen:
            return
        
        seen.add(cell)
        directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]
        for direction in directions:
            new_row, new_col = row + direction[0], col + direction[1]
            new_cell = (new_row, new_col)
            if self.is_cell_in_bounds(grid, new_row, new_col) and new_cell not in seen and grid[new_row][new_col] == '1':
                self.search(grid, seen, new_row, new_col)
    
    def is_cell_in_bounds(self, grid, row, col):
        return row >= 0 and row < len(grid) and col >= 0 and col < len(grid[0])
        
        
