class Solution:
    def cellsInRange(self, s: str) -> List[str]:
        cells = s.split(":")
        start_col, end_col = cells[0][0], cells[1][0]
        start_row, end_row = int(cells[0][1]), int(cells[1][1])
        out = []
        cur_col, cur_row = start_col, start_row

        while cur_col <= end_col:
            while cur_row <= end_row:
                out.append("{}{}".format(cur_col, cur_row))
                cur_row += 1
            cur_col, cur_row = chr(ord(cur_col) + 1), start_row

        return out
