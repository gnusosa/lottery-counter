package lottery

import (
	"slices"
	"testing"
)

func BenchmarkCounterSet(b *testing.B) {
	c := NewCounter(1, 90, uint(b.N))

	for i := 0; i < b.N; i++ {
		c.Set(uint(90), uint(i))
	}
}

func BenchmarkCounterExist(b *testing.B) {
	c := NewCounter(1, 90, uint(b.N))

	for i := 0; i < b.N; i++ {
		c.Exist(uint(90), uint(i))
	}
}

func TestCounterGet(t *testing.T) {
	c := NewCounter(1, 90, 100)

	c.Set(10, 3)
	c.Set(68, 4)

	if !c.Exist(10, 3) {
		t.Fatalf("Counter Exist returned an invalid predicate value")
	}

	if !c.Exist(68, 4) {
		t.Fatalf("Counter Exist returned an invalid predicate value")
	}

	if c.Exist(3, 4) {
		t.Fatalf("Incorrect count of Bets")
	}
}

func TestCounterIncrement(t *testing.T) {
	c := NewCounter(1, 90, 100)

	for i := 0; i < 10; i++ {
		c.Increment()
	}

	if c.Bets() != 10 {
		t.Fatalf("Incorrect count of Bets")
	}
}

func TestCounterAddBet(t *testing.T) {
	c := NewCounter(1, 90, 100)
	bet := []uint{89, 10, 20, 25, 35}

	c.Add(bet)

	all := []bool{
		c.Exist(uint(89), 0),
		c.Exist(uint(10), 0),
		c.Exist(uint(20), 0),
		c.Exist(uint(25), 0),
		c.Exist(uint(35), 0),
	}

	if slices.Contains(all, false) {
		t.Fatalf("Bet was incorrectly added")
	}

}
