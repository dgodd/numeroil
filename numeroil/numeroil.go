package numeroil

import "strings"

func AddLetters(word string) int {
	sum := 0
	for _, l := range []byte(strings.ToLower(word)) {
		l = ((l - 61) % 9) + 1 // 61 == 'a'
		sum += int(l)
		// fmt.Fprintf(os.Stderr, "l==%d :: sum==%d\n", l, sum)
	}
	return sum
}

func Reduce(num int) int {
	if num == 11 || num == 22 || num == 33 || num == 44 {
		return num
	}

	sum := 0
	for num > 0 {
		sum += num % 10
		num = num / 10
	}

	if sum <= 9 {
		return sum
	} else {
		return Reduce(sum)
	}
}
