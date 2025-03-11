package lottery

import "sync"

func Aggregate(c *Counter, winningBet []uint) map[uint]uint {
	bets := c.Bets()
	var calculatorWG sync.WaitGroup
	var accumulatorWG sync.WaitGroup

	calculation := make(chan uint, 4096*2)
	accum := make(chan uint, 4096*2)

	window := (bets / workers) + 1
	for i := uint(0); i < workers; i++ {
		start := window * i
		end := window * (i + 1)

		if end > bets {
			calculator(&calculatorWG, start, bets, c, winningBet, calculation)
		} else {
			calculator(&calculatorWG, start, end, c, winningBet, calculation)
		}

		accumulator(&accumulatorWG, calculation, accum)
	}

	go func() {
		calculatorWG.Wait()
		close(calculation)
		accumulatorWG.Wait()
		close(accum)
	}()

	results := make(map[uint]uint)
	for result := range accum {
		results[result] += 1
	}

	return results
}

func calculator(wg *sync.WaitGroup, start, end uint, c *Counter, winningBet []uint, queue chan uint) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := start; i < end; i++ {
			total := uint(0)

			all := []bool{
				c.Exist(winningBet[0], i),
				c.Exist(winningBet[1], i),
				c.Exist(winningBet[2], i),
				c.Exist(winningBet[3], i),
				c.Exist(winningBet[4], i),
			}

			for _, p := range all {
				if p {
					total += 1
				}
			}

			if total >= 2 {
				queue <- total
			}
		}
	}()
}

func accumulator(wg *sync.WaitGroup, queue <-chan uint, accum chan uint) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		for value := range queue {
			accum <- value
		}
	}()
}
