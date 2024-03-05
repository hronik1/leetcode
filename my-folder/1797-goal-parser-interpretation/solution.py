class Solution:
    def interpret(self, command: str) -> str:
        out = command
        table = {"()": "o", "(al)": "al"}
        for k, v in table.items():
            out = out.replace(k, v)
        
        return out
