# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        l = 1
        cur = head
        while cur is not None:
            cur = cur.next
            l +=1
        
        cur = head
        prev = None
        for i in range(1, l-n):
            prev = cur
            cur = cur.next
        
        if prev is None:
            return cur.next
        
        prev.next = cur.next
        
        return head
