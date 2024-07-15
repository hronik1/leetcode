class Solution:
    def findCenter(self, edges: List[List[int]]) -> int:
        seen = set()
        for edge in edges:
            for node in edge:
                if node in seen:
                    return node

                seen.add(node)

        raise ValueError("no center node")
