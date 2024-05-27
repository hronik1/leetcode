from itertools import combinations

class CombinationIterator:

    def __init__(self, characters: str, combinationLength: int):
        self.combinations = combinations(characters, combinationLength)
        self.count = math.comb(len(characters), combinationLength)
        self.i = 0

    def next(self) -> str:
        self.i += 1
        return "".join(next(self.combinations))
        
        

    def hasNext(self) -> bool:
        return self.i < self.count
        


# Your CombinationIterator object will be instantiated and called as such:
# obj = CombinationIterator(characters, combinationLength)
# param_1 = obj.next()
# param_2 = obj.hasNext()
