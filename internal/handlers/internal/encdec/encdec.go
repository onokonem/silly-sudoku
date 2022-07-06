package encdec

import (
	"crypto/rand"
	"encoding/json"
	"fmt"

	"github.com/Djarvur/cryptowrap"

	"silly-sudoku/internal/sudokutype"
)

const KeyLength = 16

var key = func() []byte { //nolint:gochecknoglobals // this is intentional
	buf := make([]byte, KeyLength)

	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}

	return buf
}()

type toPass struct {
	Secure cryptowrap.Wrapper
}

type toPassSecure[S sudokutype.Sudoku] struct {
	Field    S
	Original S
}

func newToPassSecure[S sudokutype.Sudoku](field, original S) toPassSecure[S] {
	return toPassSecure[S]{
		Field:    field,
		Original: original,
	}
}

func Enc[S sudokutype.Sudoku](field, original S) ([]byte, error) {
	srcSecure := toPassSecure[S]{Field: field, Original: original}

	src := toPass{
		Secure: cryptowrap.Wrapper{ //nolint:exhaustruct // false positive
			Keys:     [][]byte{key},
			Payload:  &srcSecure,
			Compress: true,
		},
	}

	b, err := json.Marshal(&src)
	if err != nil {
		return nil, fmt.Errorf("marshaling sudoku with cryptowrapper: %w", err)
	}

	return b, nil
}

func Dec[S sudokutype.Sudoku](data []byte, field, original S) error {
	payload := newToPassSecure(field, original)

	dst := toPass{
		Secure: cryptowrap.Wrapper{ //nolint:exhaustruct // false positive
			Keys:    [][]byte{key},
			Payload: &payload,
		},
	}

	if err := json.Unmarshal(data, &dst); err != nil {
		return fmt.Errorf("unmarshaling sudoku with cryptowrapper: %w", err)
	}

	return nil
}
