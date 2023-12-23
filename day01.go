package adventofcode2023

import "strconv"

func Day01V1(lines []string) (int, error) {
	isDigit := func(b byte) bool {
		return '0' <= b && b <= '9'
	}

	var sum int
	for _, line := range lines {
		var digits string
		for i := 0; i < len(line); i++ {
			if isDigit(line[i]) {
				digits = digits + string(line[i])
			}
		}
		if len(digits) == 1 {
			digits = string(digits[0]) + string(digits[0])
		}
		num := string(digits[0]) + string(digits[len(digits)-1])
		n, err := strconv.Atoi(num)
		if err != nil {
			return sum, err
		}
		sum = sum + n
	}
	return sum, nil
}

func Day01V2(buf []byte) (sum uint) {
	var first, last byte

	isDigit := func(b byte) bool {
		return '0' <= b && b <= '9'
	}

	for i := range buf {
		b := buf[i]
		if isDigit(buf[i]) {
			if first == 0 {
				first = b
			}
			last = b
		}
		if b == '\n' {
			sum += 10*uint(first-'0') + uint(last-'0')
			first = 0
		}
	}
	return
}
