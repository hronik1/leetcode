class Solution:
    def rearrangeArray(self, nums: List[int]) -> List[int]:
        half_len = int(len(nums)/2)
        partitioned = [0] * len(nums)
        pos_i, neg_i = 0, half_len
        for num in nums:
            if num > 0:
                partitioned[pos_i] = num
                pos_i += 1
            else:
                partitioned[neg_i] = num
                neg_i += 1

        out = []
        pos_i, neg_i = 0, half_len
        for l in range(half_len):
            out.append(partitioned[pos_i])
            out.append(partitioned[neg_i])
            pos_i, neg_i = pos_i + 1, neg_i + 1

        return out
