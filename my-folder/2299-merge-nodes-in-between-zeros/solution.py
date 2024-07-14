# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeNodes(self, head: Optional[ListNode]) -> Optional[ListNode]:
        cur = head.next
        val = cur.val
        cur = cur.next
        new_head = ListNode()
        new_prev = new_head
        while cur is not None:
            if cur.val == 0:
                new_prev.val = val
                val = 0
            elif val == 0:
                new_prev.next = ListNode()
                new_prev = new_prev.next
            
            val += cur.val
            cur = cur.next

        return new_head
