class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        rows = [[1]]
        while len(rows) < numRows:
            new_row = [1]
            prev_row = rows[-1]
            for i in range(len(prev_row)-1):
                new_row.append(prev_row[i]+prev_row[i+1])
            new_row.append(1)
            rows.append(new_row)
        return rows
