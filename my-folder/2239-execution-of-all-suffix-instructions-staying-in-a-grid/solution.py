class Solution:
    def executeInstructions(self, n: int, startPos: List[int], s: str) -> List[int]:
        return [self.steps_from_i(n, startPos, s, i) for i in range(len(s))]

    def steps_from_i(self, n, startPos, s, i):
        steps = 0
        row, col = startPos[0], startPos[1]
        directions = {'U': (-1, 0), 'D': (1, 0), 'L': (0, -1), 'R': (0, 1)}
        while i < len(s):
            direction = directions[s[i]]
            row, col = row + direction[0], col + direction[1]
            if not(row >= 0 and row < n and col >= 0 and col < n):
                break
            
            steps += 1
            i += 1
            
        return steps
