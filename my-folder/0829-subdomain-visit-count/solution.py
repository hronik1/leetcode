from collections import defaultdict

class Solution:
    def subdomainVisits(self, cpdomains: List[str]) -> List[str]:
        sum_subdomain_counts = defaultdict(int)
        for cpdomain in cpdomains:
            counts = self.subdomain_counts(cpdomain)
            for subdomain, count in counts.items():
                sum_subdomain_counts[subdomain] += count
        
        return ["{} {}".format(count, subdomain) for subdomain, count in sum_subdomain_counts.items()]

    def subdomain_counts(self, cpdomain):
        s = cpdomain.split()
        count, subdomain = s[0], s[1]
        levels = subdomain.split(".")
        return {'.'.join(levels[i:]): int(count) for i in range(len(levels))}
        
