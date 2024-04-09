from enum import Enum

class Operator(Enum):
    Open = 1
    Add = 2
    Sub = 3
    Mult = 4
    Div = 5

class Solution:
    def calculate(self, s: str) -> int:
        sanitized = self.sanitize(s)
        values = []
        op = Operator.Open
        cur_number = ''

        operator_mappings = {
            "+": Operator.Add,
            "-": Operator.Sub,
            "*": Operator.Mult,
            "/": Operator.Div,
        }

        for c in sanitized:
            if c.isnumeric():
                cur_number += c
            else:
                self.eval(cur_number, op, values)
                cur_number = ''
                op = operator_mappings[c]
            
        self.eval(cur_number, op, values)

        return sum(values)
        

    def sanitize(self, s: str) -> str:
        return s.replace(" ", "")

    def eval(self, cur_number, op, values):
        value = 0
        if cur_number != '':
            value = int(cur_number)

        if op == Operator.Open or op == Operator.Add:
            values.append(value)
        elif op == Operator.Sub:
            values.append(-1 * value)
        elif op == Operator.Mult:
            prev_value = values.pop()
            values.append(prev_value * value)
        else:
            prev_value = values.pop()
            values.append(int(prev_value / value))
        
