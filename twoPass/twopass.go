package twopass

func HowMuchSnow(mountMap []int) (int, []int) {

	// return early to avoid edge cases
	if len(mountMap) == 1 {
		return 1, []int{1}
	}

	// setup some vars
	snow := 0
	snowMap := make([]int, len(mountMap))
	i := 0
	j := len(mountMap) - 1

	// while i and j are within the bounds of mountMap, place snow
	for i < len(mountMap) && j >= 0 {

		// find the max snow for the current i index (scanning from left to right)
		iSnow := 0
		if i != 0 {
			maxSnow := snowMap[i-1] + mountMap[i-1] + 1
			if maxSnow > mountMap[i] {
				iSnow = maxSnow - mountMap[i]
			}
		} else {
			iSnow = 1
		}

		// find the max snow for the current j index (scanning from right to left)
		jSnow := 0
		if j != len(mountMap)-1 {
			maxSnow := snowMap[j+1] + mountMap[j+1] + 1
			if maxSnow > mountMap[j] {
				jSnow = maxSnow - mountMap[j]
			}
		} else {
			jSnow = 1
		}

		// before i and j meet, just place their values into the snowMap
		if i < j {
			snowMap[i] = iSnow
			snowMap[j] = jSnow
			snow += iSnow + jSnow

			// If i and j are at the same index, pick the one with the smaller values
		} else if i == j {
			if iSnow < jSnow {
				snowMap[i] = iSnow
				snow += iSnow
			} else {
				snowMap[j] = jSnow
				snow += jSnow
			}

			// after i and j meet
		} else {
			// pick the smaller value between iSnow and snowMap[i], adjust the snow total as needed
			if snowMap[i] > iSnow {
				snow = (snow - snowMap[i]) + iSnow
				snowMap[i] = iSnow
			}
			// pick the smaller value between iSnow and snowMap[i], adjust the snow total as needed
			if snowMap[j] > jSnow {
				snow -= snowMap[j]
				snow += jSnow
				snowMap[j] = jSnow
			}
		}

		// more i and j
		i = i + 1
		j = j - 1
	}
	return snow, snowMap
}
