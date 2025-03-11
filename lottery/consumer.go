package lottery

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"
)

func NewCounterFromCSVReader(r *csv.Reader) *Counter {
	c := NewCounter(minPick, maxPick, maxBets)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}

		maybeBet := make([]uint, 0, 5)
		for _, pick := range record {
			p, _ := strconv.Atoi(pick)
			maybeBet = append(maybeBet, uint(p))
		}

		if ValidBet(maybeBet) {
			c.Add(maybeBet)
		} else {
			fmt.Printf("Invalid bet found: %d", maybeBet)
			fmt.Println()
		}
	}

	return c
}
