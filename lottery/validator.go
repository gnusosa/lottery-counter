package lottery

func ValidBet(b []uint) bool {
	if len(b) < betPickLimit {
		return false
	}

	for _, p := range b {
		if !ValidPick(p) {
			return false
		}
	}

	return true
}

func ValidPick(pick uint) bool {
	if pick >= minPick && pick <= maxPick {
		return true
	}

	return false
}
