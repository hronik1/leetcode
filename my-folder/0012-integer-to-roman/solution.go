import "strings"

type Roman struct {
    V int
    S string
    SpecialV int
    SpecialS string
}

func intToRoman(num int) string {
    romans := []Roman{
        {1000, "M", -1, ""},
        {500, "D", 900, "CM"},
        {100, "C", 400, "CD"},
        {50, "L", 90, "XC",},
        {10, "X", 40, "XL",},
        {5, "V", 9, "IX"},
        {1, "I", 4, "IV"},
    }
    
    var out strings.Builder
    for _, r := range romans {
        if r.SpecialS != "" && num/r.SpecialV == 1 {
            out.WriteString(r.SpecialS)
            num = num - r.SpecialV
        } else {
            count := num/r.V 
            for i:=0; i<count; i++ {
                out.WriteString(r.S)
            }

            num = num%r.V 
        }
    }
    
    return out.String()
}
