from itertools import permutations
class Solution:
    def numTilePossibilities(self, tiles: str) -> int:
        characters = list(tiles)
        return sum([len(set(permutations(characters, i))) for i in range(1, len(characters)+1)])
