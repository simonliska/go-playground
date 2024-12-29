func convertDateToBinary(date string) string {
	parts := strings.Split(date, "-")

	y, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])
	d, _ := strconv.Atoi(parts[2])

	yToB := strconv.FormatInt(int64(y), 2)
	mToB := strconv.FormatInt(int64(m), 2)
	dToB := strconv.FormatInt(int64(d), 2)

	return fmt.Sprintf("%s-%s-%s", yToB, mToB, dToB)

}