from collections import deque
class Solution:
    def numOfSubarrays(self, arr: List[int], k: int, threshold: int) -> int:
        de = deque(arr[:k])
        avg = sum(de)/k
        count = 0
        if avg >= threshold:
            count += 1
        for i in range(k,len(arr)):
            new_item = arr[i]
            avg -= de.popleft()/k
            avg += new_item/k
            de.append(new_item)
            if avg >= threshold:
                count += 1
        
        return count
        
