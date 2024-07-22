class Solution:
    def onesMinusZeros(self, grid: List[List[int]]) -> List[List[int]]:
        onesRow = [sum(row) for row in grid]
        zerosRow = [len(grid) - oneRow for oneRow in onesRow]
        onesCol= [sum([row[j] for row in grid]) for j in range(len(grid[0]))]
        zerosCol = [len(grid[0]) - oneCol for oneCol in onesCol]

        out = [ [0]*len(grid[0]) for i in range(len(grid))]
        print(onesRow[0] + onesCol[0] - zerosRow[0] - zerosCol[0])
        for i in range(len(grid)):
            for j in range(len(grid[0])):
                s = onesRow[i] + onesCol[j] - zerosRow[i] - zerosCol[j]
                out[i][j] = s
        
        return out
