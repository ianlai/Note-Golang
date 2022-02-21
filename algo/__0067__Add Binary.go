func addBinary(a string, b string) string {
	ra := reverse(a)
	rb := reverse(b)
	maxLen := max(len(a), len(b))
	//res := make([]byte, maxLen)
	carry := 0
	var buffer bytes.Buffer
	for i := 0; i < maxLen; i++ {
		sum := carry
		val := 0
		if i < len(ra) {
			sum += getInt(i, ra)
		}
		if i < len(rb) {
			sum += getInt(i, rb)
		}
		if sum == 0 {
			carry = 0
			val = 0
			//res = append(res, 0)
		} else if sum == 1 {
			carry = 0
			val = 1
			//res = append(res, 1)
		} else if sum == 2 {
			carry = 1
			val = 0
			//res = append(res, 0)
		} else {
			carry = 1
			val = 1
			//res = append(res, 1)
		}
		buffer.WriteString(strconv.Itoa(val))
	}
	if carry == 1 {
		buffer.WriteString(strconv.Itoa(1))
	}
	//return string(res)
	return reverse(buffer.String())
}
func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func getInt(index int, s string) int {
	part, _ := strconv.Atoi(string(s[index]))
	return part
}