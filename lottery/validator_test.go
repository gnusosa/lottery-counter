package lottery

import "testing"

func TestValidBet(t *testing.T) {
	bet := make([]uint, 5)

	bet[0] = 11
	bet[1] = 12
	bet[2] = 13
	bet[3] = 14
	bet[4] = 15

	if !ValidBet(bet) {
		t.Fatalf("ValidBet predicate function should return true")
	}
}

func TestValidBetInvalid(t *testing.T) {
	bet := make([]uint, 5)

	bet[0] = 11
	bet[1] = 12
	bet[2] = 13
	bet[3] = 14
	bet[4] = 95

	if ValidBet(bet) {
		t.Fatalf("ValidBet predicate function should return false")
	}
}

func TestValidBetInvalidLen(t *testing.T) {
	bet := make([]uint, 5)

	bet[0] = 11
	bet[1] = 12
	bet[2] = 13
	bet[3] = 14

	if ValidBet(bet) {
		t.Fatalf("ValidBet predicate function should return false")
	}
}
