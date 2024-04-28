package adventofcode2023

import "fmt"

func Day03(lines []string) (sum int) {
	var readNumber bool
	var n, l int
	for y := range lines {
		for x := range lines[y] {
			c := lines[y][x]
			if isDigit(c) {
				readNumber = true
				n = 10*n + int(c-'0')
				l++
			} else {
				if readNumber {
					// border contains symbol?
					surrounded := false
					x1 := max(x-l-1, 0)
					x2 := min(x, len(lines[y])-1)

					// top row
					if y > 0 {
						for i := x1; i <= x2; i++ {
							if isSymbol(lines[y-1][i]) {
								surrounded = true
								goto done
							}
						}
					}

					// left/ right
					if isSymbol(lines[y][x1]) || isSymbol(lines[y][x2]) {
						surrounded = true
						goto done
					}

					// bottom row
					if y+1 < len(lines) {
						for i := x1; i <= x2; i++ {
							if isSymbol(lines[y+1][i]) {
								surrounded = true
								goto done
							}
						}
					}
				done:
					if surrounded {
						fmt.Printf("adding %d\n", n)
						sum += n
					}
					// done reading number
					n, l = 0, 0
					readNumber = false
					surrounded = false
				}
			}
		}
	}
	return
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isSymbol(b byte) bool {
	return !(b == '.' || isDigit(b))
}
