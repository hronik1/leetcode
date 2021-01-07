import "fmt"

func defangIPaddr(address string) string {
    out := ""
    
    for _, r := range address {
        if string(r) == "." {
            out += fmt.Sprintf("[%s]", string(r))
        } else {
            out += string(r)
        }
    }
    
    return out
}
