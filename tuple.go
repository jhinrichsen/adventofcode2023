package adventofcode2023

import (
	"fmt"
)

type Tuple[E any] struct {
	A, B E
}

func (t Tuple[E]) Len() int {
	return 2
}

func (t Tuple[E]) String() string {
	return fmt.Sprintf("(%v/%v)", t.A, t.B)
}
