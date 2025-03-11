package lottery

import "github.com/bits-and-blooms/bitset"

type Counter struct {
	bets uint
	s    []*bitset.BitSet
}

func NewCounter(minPick, maxPick, maxBets uint) *Counter {
	c := &Counter{
		s: make([]*bitset.BitSet, maxPick+1),
	}

	for i := minPick; i <= maxPick; i++ {
		c.s[i] = bitset.New(maxBets)
	}

	return c
}

func (c *Counter) Set(pick, bet uint) {
	(c.s[pick]).Set(bet)
}

func (c *Counter) Exist(pick, bet uint) bool {
	return (c.s[pick]).Test(bet)
}

func (c *Counter) Increment() {
	c.bets += 1
}

func (c *Counter) Bets() uint {
	return c.bets
}

func (c *Counter) Add(bet []uint) {
	for _, pick := range bet {
		c.Set(pick, c.Bets())
	}

	c.Increment()
}
