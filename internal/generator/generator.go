package generator

import (
	"math/rand"
	"time"

	"silly-sudoku/internal/solver"
	"silly-sudoku/internal/sudokutype"
)

// Generate is filling the provided sudoku with tha random generated one.
// Original data is ignored and overwritten.
func Generate[S sudokutype.Sudoku](s S) {
	sudokutype.BaseFill(s)
	Randomize(s, rand.New(rand.NewSource(time.Now().UnixNano()))) //nolint:gosec // math/rand is ok
}

// Suppress is zeroing the cells with the provided probability.
func Suppress[S sudokutype.Sudoku](in, out, tmp S, difficulty float64) S { //nolint:ireturn //false positive
	rnd := rand.New(rand.NewSource(time.Now().UnixNano())) //nolint:gosec // math/rand is ok

	for {
		sudokutype.Copy(out, in)

		for y := 0; y < len(out); y++ {
			for x := 0; x < len(out); x++ {
				if rnd.Float64() > difficulty {
					out.Set(y, x, 0)
				}
			}
		}

		sudokutype.Copy(tmp, out)

		if ok, _ := solver.Solve(tmp); ok {
			return out
		}
	}
}
