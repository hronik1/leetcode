class Solution:
    def minOperations(self, boxes: str) -> List[int]:
        balls = set([i for i in range(len(boxes)) if boxes[i] == '1'])
        
        out = []
        for i in range(len(boxes)):
            s = sum([abs(i-ball) for ball in balls])
            out.append(s)
        return out 
