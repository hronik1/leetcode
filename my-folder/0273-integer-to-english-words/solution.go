var singles = map[int]string {
    0: "Zero",
    1: "One",
    2: "Two",
    3: "Three",
    4: "Four",
    5: "Five",
    6: "Six",
    7: "Seven",
    8: "Eight",
    9: "Nine",
}

var tens = map[int]string {
    10: "Ten",
    11: "Eleven",
    12: "Twelve",
    13: "Thirteen",
    14: "Fourteen",
    15: "Fifteen",
    16: "Sixteen",
    17: "Seventeen",
    18: "Eighteen",
    19: "Nineteen",
    20: "Twenty",
    30: "Thirty",
    40: "Forty",
    50: "Fifty",
    60: "Sixty",
    70: "Seventy",
    80: "Eighty",
    90: "Ninety",
}

var hundred = "Hundred"
var thousand = "Thousand"
var million = "Million"
var billion = "Billion"

func numberToWords(num int) string {
    out := []string{}
    
    if num == 0 {
        out = append(out, singles[num])
    }
    
    for num > 0 {
        if num >= 1000000000 {
            out = append(out, thousandsNumToWord(num/1000000000)...)
            out = append(out, billion)
            num %= 1000000000
        } else if num >= 1000000 {
            out = append(out, thousandsNumToWord(num/1000000)...)
            out = append(out, million)
            num %= 1000000
        } else if num >= 1000 { 
            out = append(out, thousandsNumToWord(num/1000)...)
            out = append(out, thousand)
            num %= 1000
        } else {
            out = append(out, thousandsNumToWord(num)...)
            break
        }
        
    }
    
    return strings.Join(out, " ")
}

func thousandsNumToWord(thousandNum int) []string {
    out := []string{}
    for thousandNum > 0 {
        if thousandNum >= 100 {
            out = append(out, singles[thousandNum/100])
            out = append(out, hundred)
            thousandNum %= 100
        } else if thousandNum >= 90 {
            out = append(out, tens[90])
            thousandNum %= 90
        } else if thousandNum >= 80 {
            out = append(out, tens[80])
            thousandNum %= 80
        } else if thousandNum >= 70 {
            out = append(out, tens[70])
            thousandNum %= 70 
        } else if thousandNum >= 60 {
            out = append(out, tens[60])
            thousandNum %= 60    
        } else if thousandNum >= 50 {
            out = append(out, tens[50])
            thousandNum %= 50            
        } else if thousandNum >= 40 {
            out = append(out, tens[40])
            thousandNum %= 40
        } else if thousandNum >= 30 {
            out = append(out, tens[30])
            thousandNum %= 30
        } else if thousandNum >= 20 {
            out = append(out, tens[20])
            thousandNum %= 20            
        } else if thousandNum >= 10 {
            out = append(out, tens[thousandNum])
            break
        } else {
            out = append(out, singles[thousandNum])
            break
        }
        
    }
    
    return out
} 


