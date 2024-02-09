class Solution:
    def validateStackSequences(self, pushed: List[int], popped: List[int]) -> bool:
        pushedI, poppedI = 0, 0
        stack = []
        while poppedI < len(popped):
            if len(stack) > 0 and stack[-1] == popped[poppedI]:
                stack.pop()
                poppedI += 1
            elif pushedI < len(pushed):
                stack.append(pushed[pushedI])
                pushedI += 1
            else:
                return False

        return True
        
        
