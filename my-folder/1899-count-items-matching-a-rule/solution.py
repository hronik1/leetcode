class Solution:
    def countMatches(self, items: List[List[str]], ruleKey: str, ruleValue: str) -> int:
        index = 0
        if ruleKey == "color":
            index = 1
        elif ruleKey == "name":
            index = 2
        
        count = 0
        for item in items:
            if item[index] == ruleValue: count += 1
        
        return count
            
