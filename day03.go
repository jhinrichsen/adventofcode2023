package adventofcode2023

import (
	"fmt"

	"github.com/fatih/color"
)

type Day03Puzzle []string

func NewDay03(lines []string) (Day03Puzzle, error) {
	return Day03Puzzle(lines), nil
}

func Day03(puzzle Day03Puzzle, part1 bool) int {
	if part1 {
		return Day03Part1([]string(puzzle))
	}
	return int(Day03Part2([]string(puzzle)))
}

func Day03Part1(lines []string) (sum int) {
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
	enter := func() {
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
					enter()
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

func Day03Part2(lines []string) (sum uint) {
	const maxNeighbours = 6

	var numbers [maxNeighbours]uint // avoid heap allocs
	var n byte                      // index into numbers

	maxX := len(lines[0]) - 1
	maxY := len(lines) - 1

	// 'A gear is any * symbol that is adjacent to exactly two part numbers.'
	isGear := func() bool {
		return n == 2
	}

	gearRatio := func() uint {
		if !isGear() {
			return 0
		}
		return numbers[0] * numbers[1]
	}

	left := func(x, y int) int {
		for x > 0 && isDigit(lines[y][x]) {
			x--
		}
		return x
	}
	right := func(x, y int) int {
		for x < maxX && isDigit(lines[y][x]) {
			x++
		}
		return x
	}
	readNumbers := func(x1, x2, y int) {
		// parse number, optionally multiple numbers separated by space
		var val uint
		store := func() {
			numbers[n] = val
			n++
			val = 0
		}

		// trim
		for x1 < x2 && !isDigit(lines[y][x1]) {
			x1++
		}
		for x1 < x2 && !isDigit(lines[y][x2]) {
			x2--
		}
		// found a number after trimming?
		if x1 == x2 && !isDigit(lines[y][x1]) {
			return
		}
		for i := x1; i <= x2; i++ {
			if !isDigit(lines[y][i]) {
				store()
			} else {
				val = 10*val + digit(lines[y][i])
			}
		}
		store()
	}

	for y := range lines {
		for x := range maxX {
			if !isGearSymbol(lines[y][x]) {
				continue
			}

			n = 0
			// North
			// TODO remove safe, backed min/max combos
			if y-1 >= 0 {
				readNumbers(left(max(0, x-1), y-1),
					right(min(maxX, x+1), y-1), y-1)
			}

			// West
			if x-1 >= 0 {
				readNumbers(left(max(0, x-1), y), max(0, x-1), y)
			}

			// East
			if x+1 <= maxX {
				readNumbers(min(maxX, x+1), right(min(maxX, x+1), y), y)
			}

			// South
			if y+1 <= maxY {
				readNumbers(left(max(0, x-1), y+1),
					right(min(maxX, x+1), y+1), y+1)
			}

			sum += gearRatio()
		}
	}
	return
}

func digit(b byte) uint {
	return uint(b - '0')
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isGearSymbol(b byte) bool {
	return b == '*'
}

func isSpace(b byte) bool {
	return b == '.'
}

func isSymbol(b byte) bool {
	return !(isSpace(b) || isDigit(b))
}
