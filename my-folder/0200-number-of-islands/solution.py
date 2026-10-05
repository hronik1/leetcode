class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        if not len(grid) or not len(grid[0]):
            return 0

        num_islands = 0
        visited = set()
        for i in range(len(grid)):
            for j in range(len(grid[0])):
                if grid[i][j] == "1" and (i, j) not in visited:
                    num_islands += 1
                    visited.add((i, j))
                    self.visit(grid, i+1, j, visited)
                    self.visit(grid, i, j+1, visited)
        
        return num_islands

    def visit(self, grid: List[List[str]], i: int, j: int, visited: set[tuple[int, int]]) -> None:
        if i >= len(grid) or i < 0 or j >= len(grid[i]) or j < 0:
            return
        
        if grid[i][j] == "0" or (i, j) in visited:
            return

        visited.add((i, j))
        self.visit(grid, i+1, j, visited)
        self.visit(grid, i, j+1, visited)
        self.visit(grid, i-1, j, visited)
        self.visit(grid, i, j-1, visited)
        
