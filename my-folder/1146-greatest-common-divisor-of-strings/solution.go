func gcdOfStrings(str1 string, str2 string) string {
    if str1 + str2 == str2 + str1 {
        return str1[:gcd(len(str1), len(str2))]
    } else {
        return ""
    }
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
