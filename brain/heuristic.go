package player

var undefinedTile = tile{255, 255}

const (
	commonRatio = 3
	startWeight = 10000000
)

func eval(b uint64) (int, tile) {
	maxLinearWeight := -1
	critical := undefinedTile

	for _, p := range paths {
		criticalTile := undefinedTile
		weight := startWeight
		linearWeight := 0
		for _, t := range p {
			v := get(b, t.y, t.x)
			if v == 0 && criticalTile == undefinedTile {
				criticalTile = tile{t.y, t.x}
			}
			linearWeight += pow2[v] * weight
			weight /= commonRatio
		}
		if linearWeight > maxLinearWeight {
			maxLinearWeight = linearWeight
			critical = criticalTile
		}
	}

	return maxLinearWeight, critical
}
