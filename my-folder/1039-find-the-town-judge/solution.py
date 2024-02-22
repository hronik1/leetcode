class Solution:
    def findJudge(self, n: int, trust: List[List[int]]) -> int:
        not_truster = set([i for i in range(1, n+1)])
        trusted_count = defaultdict(int)

        for edge in trust:
            if edge[0] in not_truster:
                not_truster.remove(edge[0])
            trusted_count[edge[1]] += 1
        
        if len(not_truster) != 1:
            return -1
        
        potential_judge = list(not_truster)[0]
        if trusted_count[potential_judge] != n - 1:
            return -1
        
        return potential_judge
