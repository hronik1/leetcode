class Solution:
    def isValidSudoku(self, board: list[list[str]]) -> bool:
        if not self.areRowsValid(board):
            return False
        
        if not self.areColumnsValid(board):
            return False
        
        if not self.areBoxesValid(board):
            return False
        
        return True
    
    def areRowsValid(self, board: list[list[str]]) -> bool:
        for i in range(9):
            if not self.isRowValid(board[i]):
                return False
        
        return True
    
    def isRowValid(self, row: list[str]) -> bool:
        seen = set()
        for c in row:
            if c == ".":
                continue
            if c in seen:
                return False
            seen.add(c)
        
        return True

    def areColumnsValid(self, board: list[list[str]]) -> bool:
        for i in range(9):
            if not self.isColumnValid(board, i):
                return False
        
        return True
    
    def isColumnValid(self, board: list[list[str]], column: int) -> bool:
        seen = set()
        for i in range(9):
            c = board[i][column]
            if c == ".":
                continue

            if c in seen:
                return False
            seen.add(c)
        
        return True         
    
    def areBoxesValid(self, board: list[list[str]]) -> bool:
        for i in [0, 3, 6]:
            for j in [0, 3, 6]:
                if not self.isBoxValid(board, i, j):
                    return False

        return True
    
    def isBoxValid(self, board: list[list[str]], row_i: int, col_j: int) -> bool:
        end_i, end_j = row_i + 3, col_j + 3
        seen = set()
        for i in range(row_i, end_i):
            for j in range(col_j, end_j):
                c = board[i][j]
                if c == ".":
                    continue
                    
                if c in seen:
                    return False
                seen.add(c)
        
        return True
