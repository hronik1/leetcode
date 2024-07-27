class Solution:
    def findSmallestSetOfVertices(self, n: int, edges: List[List[int]]) -> List[int]:
        possible_nodes = set(range(n))
        for edge in edges:
            if edge[1] in possible_nodes:
                possible_nodes.remove(edge[1])

        return list(possible_nodes)
