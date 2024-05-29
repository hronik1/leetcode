# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def insertGreatestCommonDivisors(self, head: Optional[ListNode]) -> Optional[ListNode]:
        cur, nxt = head, head.next
        while nxt is not None:
            gcd = math.gcd(cur.val, nxt.val)
            node = ListNode(gcd, nxt)
            cur.next = node
            cur, nxt = nxt, nxt.next
        
        return head
