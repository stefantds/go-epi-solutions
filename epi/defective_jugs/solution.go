package defective_jugs

type volumeRange struct {
	low, high int
}

func CheckFeasible(jugs []Jug, l int, h int) bool {
	cache := make(map[volumeRange]bool)
	return checkFeasibleWithCache(jugs, l, h, cache)
}

func checkFeasibleWithCache(jugs []Jug, l int, h int, cacheUnsuccessful map[volumeRange]bool) bool {
	if l < 0 || h < 0 || l > h {
		return false
	}
	if cacheUnsuccessful[volumeRange{low: l, high: h}] {
		return false
	}

	for _, j := range jugs {
		if jugSatisfiesConstraint(j, l, h) {
			return true
		}
		if checkFeasibleWithCache(jugs, l-j.Low, h-j.High, cacheUnsuccessful) {
			return true
		}
	}
	cacheUnsuccessful[volumeRange{low: l, high: h}] = true
	return false
}

func jugSatisfiesConstraint(jug Jug, low, high int) bool {
	return jug.High <= high && jug.Low >= low
}
