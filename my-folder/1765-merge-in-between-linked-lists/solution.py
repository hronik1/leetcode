# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeInBetween(self, list1: ListNode, a: int, b: int, list2: ListNode) -> ListNode:
        beforeA, afterB = None, None
        cur = list1
        i = 0
        while cur is not None and i <= b:
            if i+1 == a:
                beforeA = cur
            if i == b:
                afterB = cur.next

            cur = cur.next
            i += 1
        
        cur2 = list2
        while cur2.next is not None:
            cur2 = cur2.next
        
        beforeA.next = list2
        cur2.next = afterB

        return list1
        
