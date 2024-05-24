from collections import Counter
class Solution:
    def countValidWords(self, sentence: str) -> int:
        tokens = self.tokenize(sentence)
        return len([token for token in tokens if self.is_valid_token(token)])

    def tokenize(self, sentence):
        return sentence.split()
    
    def is_valid_token(self, token):
        counter = Counter(token)
        hyphen_count = counter['-']
        if hyphen_count > 1:
            return False
        
        if hyphen_count == 1:
            hyphen_i = token.find('-')
            if hyphen_i == 0 or hyphen_i == len(token)-1:
                return False
            
            prev_c, next_c = token[hyphen_i-1], token[hyphen_i+1]
            if not prev_c.islower() or not next_c.islower():
                return False
        
        punctuation = set(['!', ',', '.'])
        punctuation_count = sum([counter[c] for c in punctuation])
        if punctuation_count > 1:
            return False
        
        if punctuation_count == 1 and token[-1] not in punctuation:
            return False
        
        digit_count = sum([counter['{}'.format(digit)] for digit in range(10)])
        if digit_count != 0:
            return False

        return True
