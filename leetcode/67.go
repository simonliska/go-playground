func addBinary(a string, b string) string {
	decimalA := new(big.Int)
	decimalB := new(big.Int)

	_, successA := decimalA.SetString(a, 2)
	_, successB := decimalB.SetString(b, 2)

	if !successA || !successB {
		return ""
	}

	sum := new(big.Int).Add(decimalA, decimalB)

	return sum.Text(2)

}