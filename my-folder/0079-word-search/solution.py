class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        visited = set()
        for row in range(len(board)):
            for col in range(len(board[row])):
                if self.existsHelper(board, word, row, col, 0, visited):
                    return True
        
        return False


    def existsHelper(self, board: List[List[str]], word: str, row: int, col: int, i: int, visited: set[tuple[int, int]]) -> bool:
        if word[i] != board[row][col]:
            return False

        if i == len(word) - 1:
            return True

        visited.add((row, col))
        directions = [[-1, 0], [1, 0], [0, -1], [0, 1]]
        for direction in directions:
            newRow, newCol = row + direction[0], col + direction[1]
            if self.isInBounds(board, newRow, newCol) and (newRow, newCol) not in visited:
                if self.existsHelper(board, word, newRow, newCol, i+1, visited):
                    return True


        visited.remove((row, col))
        return False

    def isInBounds(self, board: List[List[str]], row: int, col: int) -> bool:
        return row >= 0 and col >= 0 and row < len(board) and col < len(board[row])
