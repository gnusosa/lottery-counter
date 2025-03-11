package lottery

import (
	"bufio"
	"fmt"
	"gnusosa/lottery-counter/bet"
	"log"
	"os"
	"strconv"
	"time"
)

func Execute() {
	// get data file
	f := GetArgsFile()

	// // create counter from file stream
	br := bet.NewReader(f)
	c := NewCounterFromCSVReader(br)

	var err error
	if err = f.Close(); err != nil {
		panic(err)
	}

	// consume winning bet from STDIN
	var winningBet []uint
	winningBet, err = GetWinningBet()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Winning bet:")
	fmt.Println(winningBet)

	start := time.Now()
	results := Aggregate(c, winningBet)
	end := time.Since(start)

	log.Printf("took: %d ms", end.Milliseconds())
	log.Println()

	for i := uint(5); i >= 2; i-- {
		fmt.Printf("%d: %d\n", i, results[i])
	}
}

func GetArgsFile() *os.File {
	file := os.Args[1:2][0]

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	return f
}

func GetWinningBet() ([]uint, error) {
	br := bet.NewReader(bufio.NewReader(os.Stdin))
	fmt.Println("READY")

	maybeBet := make([]uint, 5)
	bet, err := br.Read()
	if err != nil {
		log.Fatal(err)
	}

	for _, pick := range bet {
		maybePick, err := strconv.Atoi(pick)
		if err != nil {
			return []uint{}, fmt.Errorf("skipping incorrect pick: %d. Error msg: %s", maybePick, err)
		}

		maybeBet = append(maybeBet, uint(maybePick))
	}

	maybeBet = maybeBet[5:]
	if ValidBet(maybeBet) {
		return maybeBet, nil
	}

	return []uint{}, fmt.Errorf("incorrect bet input: %d", maybeBet)
}
