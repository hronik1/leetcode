class Solution:
    def numberOfBeams(self, bank: List[str]) -> int:
        beams = 0
        prev_count = 0
        for row in bank:
            cur_count = row.count('1')
            if cur_count > 0:
                beams += cur_count * prev_count
                prev_count = cur_count

        return beams
