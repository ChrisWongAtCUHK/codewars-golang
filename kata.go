package kata

// count the number of bit '1'
func countBit(n int) int {
	cnt := 1
	for n > 1 {
		if n % 2 == 1 {
			cnt += 1
		}
		n = n / 2
	}
	return cnt
}

func NextHigher(n int) int {
	bitCnt := countBit(n)
	n += 1

	for bitCnt != countBit(n) {
		n += 1
	}
	return n
}
