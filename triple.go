package adventofcode2023

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
