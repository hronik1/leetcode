class Solution:
    def fillCups(self, amount: List[int]) -> int:
        sorted_amounts = sorted(amount)
        fills = 0
        while sorted_amounts[1] > 0:
            fills += 1
            sorted_amounts[1] -= 1
            sorted_amounts[2] -= 1
            sorted_amounts = sorted(sorted_amounts)

        return fills + sorted_amounts[2]
