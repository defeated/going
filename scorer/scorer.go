package scorer

import (
	"time"
)

// from http://wiki.mozilla.org/User:Mconnor/PlacesFrecency#Option_2_.28lowest_overhead.2Fgood_accuracy.29
// 1. Each hit is worth 100 points
// 2. Divide the total points by the days since first visit
// 3. Round to nearest 10
func FrecencyScore(h uint16, t time.Time) uint16 {
	points := h * 100
	days := time.Since(t).Hours() / 24
	score := points / uint16(days)
	// TODO: round
	return score
}
