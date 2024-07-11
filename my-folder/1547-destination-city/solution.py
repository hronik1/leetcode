class Solution:
    def destCity(self, paths: List[List[str]]) -> str:
        sources, destinations = set([path[0] for path in paths]), set([path[1] for path in paths])
        return list(destinations - sources)[0]
