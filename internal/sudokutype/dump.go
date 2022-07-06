package sudokutype

import (
	"bytes"
	"fmt"
	"strconv"
)

// Dump returns a sudoku printed as a table.
func Dump[S Sudoku](s S) string {
	var res bytes.Buffer

	f := "%" + strconv.Itoa(len(strconv.Itoa(len(s)))) + "s|"

	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s); x++ {
			if v := s.Get(y, x); v != 0 {
				fmt.Fprintf(&res, f, strconv.Itoa(int(v)))
			} else {
				fmt.Fprintf(&res, f, "")
			}
		}
		fmt.Fprintf(&res, "\n")
	}

	return res.String()
}
