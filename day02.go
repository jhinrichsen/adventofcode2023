package adventofcode2023

type Day02Puzzle []string

func NewDay02(lines []string) (Day02Puzzle, error) {
	return Day02Puzzle(lines), nil
}

func Day02(puzzle Day02Puzzle, part1 bool) uint {
	var sum uint
	ref := Triple{12, 13, 14}

	for _, line := range puzzle {
		gameID, maxTriple := parseDay02Line(line)

		if part1 {
			if maxTriple.Within(ref) {
				sum += gameID
			}
		} else {
			sum += maxTriple.Power()
		}
	}
	return sum
}

func parseDay02Line(line string) (uint, Triple) {
	// Format: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	var gameID uint
	var maxTriple Triple

	i := 5 // skip "Game "
	// Parse game ID
	for i < len(line) && line[i] >= '0' && line[i] <= '9' {
		gameID = gameID*10 + uint(line[i]-'0')
		i++
	}
	i += 2 // skip ": "

	// Parse color reveals
	var n uint
	colorStart := 0

	for i < len(line) {
		c := line[i]

		if c >= '0' && c <= '9' {
			n = n*10 + uint(c-'0')
		} else if c == ' ' && n > 0 {
			// Space after number - color name starts next
			colorStart = i + 1
		} else if c == ',' || c == ';' || i == len(line)-1 {
			// End of color entry
			if i == len(line)-1 && c != ',' && c != ';' {
				// Include last character if it's part of color name
				i++
			}

			if n > 0 && colorStart > 0 {
				// Determine color by first character (r/g/b are unique)
				switch line[colorStart] {
				case 'r': // red
					maxTriple.A = max(maxTriple.A, n)
				case 'g': // green
					maxTriple.B = max(maxTriple.B, n)
				case 'b': // blue
					maxTriple.C = max(maxTriple.C, n)
				}
			}
			n = 0
			colorStart = 0
		}
		i++
	}

	return gameID, maxTriple
}
