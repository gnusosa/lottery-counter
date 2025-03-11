package bet

import (
	"encoding/csv"
	"io"
)

func NewReader(r io.Reader) *csv.Reader {
	br := csv.NewReader(r)
	br.Comma = ' '

	return br
}
