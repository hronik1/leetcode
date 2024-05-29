class Solution:
    def maxIceCream(self, costs: List[int], coins: int) -> int:
        sorted_costs = sorted(costs)
        count, spend = 0, 0
        for sorted_cost in sorted_costs:
            spend += sorted_cost
            if spend > coins:
                break
            
            count += 1
        
        return count
