class Solution:
    def findOrder(self, numCourses: int, prerequisites: List[List[int]]) -> List[int]:
        out = []
        inboundPrereqCounts = dict([(i, 0) for i in range(numCourses)])
        outboundPrereqs = dict([(i, set()) for i in range(numCourses)])
        
        for prereq in prerequisites:
            inboundPrereqCounts[prereq[0]] += 1
            outboundPrereqs[prereq[1]].add(prereq[0])


        noPrereqs = set([i for i, count in inboundPrereqCounts.items() if count == 0])
        while len(noPrereqs) > 0:
            newNoPrereqs = set()
            for noPrereq in noPrereqs:
                out.append(noPrereq)
                for outboundPrereq in outboundPrereqs[noPrereq]:
                    inboundPrereqCounts[outboundPrereq] -= 1
                    if inboundPrereqCounts[outboundPrereq] == 0:
                        newNoPrereqs.add(outboundPrereq)

            noPrereqs = newNoPrereqs
            


        if len(out) != numCourses:
            return []

        return out
        
