class Solution:
    def countPoints(self, points: List[List[int]], queries: List[List[int]]) -> List[int]:
        out = []
        for query in queries:
            out.append(sum([1 if self.is_in_query(point, query) else 0 for point in points]))
        return out
    
    def is_in_query(self, point, query):
        return math.dist(point, query[:2]) <= query[2]
