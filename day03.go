package adventofcode2023

import (
	"fmt"

	"github.com/fatih/color"
)

func Day03(lines []string) (sum int) {
	isSurrounded := func(x, y, l int) bool {
		x1 := max(x-l-1, 0)
		x2 := min(x, len(lines[y])-1)

		// top row
		if y > 0 {
			for i := x1; i <= x2; i++ {
				if isSymbol(lines[y-1][i]) {
					return true
				}
			}
		}

		// left/ right
		if isSymbol(lines[y][x1]) || isSymbol(lines[y][x2]) {
			return true
		}

		// bottom row
		if y+1 < len(lines) {
			for i := x1; i <= x2; i++ {
				if isSymbol(lines[y+1][i]) {
					return true
				}
			}
		}
		return false
	}

	// events for digits and numbers

	var n, l int
	exit := func(x, y int) {
		if isSurrounded(x, y, l) {
			sum += n
		}
	}
	digit := func(b byte) {
		n = 10*n + int(b-'0')
		l++
	}

	for y := range lines {
		high, last := false, false
		for x := range lines[y] {
			c := lines[y][x]
			if isDigit(c) {
				high = true
				if high != last {
					n, l = 0, 0
				}
				last = high
				digit(c)
			} else {
				high = false
				if high != last {
					exit(x, y)
				}
				last = high
			}
		}
		if high {
			exit(len(lines[y]), y)
		}
	}
	return
}
func Day03ColoredLogging(lines []string) (sum int) {
	isSurrounded := func(x, y, l int) bool {
		x1 := max(x-l-1, 0)
		x2 := min(x, len(lines[y])-1)

		// top row
		if y > 0 {
			for i := x1; i <= x2; i++ {
				if isSymbol(lines[y-1][i]) {
					return true
				}
			}
		}

		// left/ right
		if isSymbol(lines[y][x1]) || isSymbol(lines[y][x2]) {
			return true
		}

		// bottom row
		if y+1 < len(lines) {
			for i := x1; i <= x2; i++ {
				if isSymbol(lines[y+1][i]) {
					return true
				}
			}
		}
		return false
	}

	// events for digits and numbers

	var n, l int
	enter := func(x, y int) {
		n, l = 0, 0
		// color.Set(color.FgYellow)
	}
	exit := func(x, y int) {
		if isSurrounded(x, y, l) {
			sum += n
			color.Set(color.FgGreen)
		} else {
			color.Set(color.FgRed)
		}
		fmt.Printf("%d", n)
		color.Unset()
	}
	digit := func(b byte) {
		n = 10*n + int(b-'0')
		l++
	}

	for y := range lines {
		high, last := false, false
		for x := range lines[y] {
			c := lines[y][x]
			if isDigit(c) {
				high = true
				if high != last {
					enter(x, y)
				}
				last = high
				digit(c)
			} else {
				high = false
				if high != last {
					exit(x, y)
				}
				last = high
			}
			if !high {
				fmt.Printf("%c", lines[y][x])
			}
		}
		if high {
			exit(len(lines[y]), y)
		}
		fmt.Println()
	}
	return
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isSymbol(b byte) bool {
	return !(b == '.' || isDigit(b))
}
