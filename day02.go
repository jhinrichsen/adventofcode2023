package adventofcode2023

import (
	"fmt"
	"strconv"
	"strings"
)

type Triple struct {
	A, B, C uint
}

func (a Triple) Within(t Triple) bool {
	return a.A <= t.A && a.B <= t.B && a.C <= t.C
}

func (a Triple) Power() uint {
	return a.A * a.B * a.C
}

func Max(a, b Triple) Triple {
	return Triple{
		max(a.A, b.A),
		max(a.B, b.B),
		max(a.C, b.C),
	}
}

func Day02(ref Triple, lines []string, part1 bool) (uint, error) {
	var sum uint

	if !part1 {
		ref = Triple{}
	}
	for i, line := range lines {
		game, t, err := parseDay02Line(line)
		if err != nil {
			return 0, fmt.Errorf("error parsing line %d: %v", i+1, err)
		}
		if part1 {
			if t.Within(ref) {
				sum += game
			}
		} else {
			sum += t.Power()
		}
	}
	return sum, nil
}

// returns game ID, r, g, b values
// format: Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// colors have no fixed position
func parseDay02Line(s string) (game uint, t Triple, err error) {
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		err = fmt.Errorf("want two parts but got %d", len(parts))
		return
	}
	var i int
	i, err = strconv.Atoi(parts[0][5:])
	if err != nil {
		return
	}
	game = uint(i)
	turns := strings.Split(parts[1], ";")
	for _, turn := range turns {
		cols := strings.Split(turn, ",")
		for _, col := range cols {
			colvals := strings.Fields(col)
			if len(colvals) != 2 {
				err = fmt.Errorf("want <n> <color> but got %q", col)
				return
			}
			var i int
			i, err = strconv.Atoi(colvals[0])
			if err != nil {
				err = fmt.Errorf("error parsing %q: %v", colvals, err)
				return
			}
			if i < 0 {
				panic(fmt.Sprintf("want i >= 0 but got %d", i))
			}
			n := uint(i)
			// cannot have f func of [min|max] because they are builtins
			switch colvals[1] {
			case "red":
				t.A = max(t.A, n)
			case "green":
				t.B = max(t.B, n)
			case "blue":
				t.C = max(t.C, n)
			}
		}
	}
	return
}
