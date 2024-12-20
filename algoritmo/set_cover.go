package algoritmo

func MinimumSetCover(universalSet []int, sets [][]int) [][]int {
	remaining := make(map[int]bool)
	for _, element := range universalSet {
		remaining[element] = true
	}

	selectedSets := [][]int{}

	for len(remaining) > 0 {
		bestSet := []int{}
		covered := 0

		for _, set := range sets {
			currentCovered := 0
			for _, element := range set {
				if remaining[element] {
					currentCovered++
				}
			}

			if currentCovered > covered {
				bestSet = set
				covered = currentCovered
			}
		}

		if len(bestSet) == 0 {
			break
		}

		selectedSets = append(selectedSets, bestSet)

		for _, element := range bestSet {
			delete(remaining, element)
		}
	}

	return selectedSets
}
