package multipass

import (
	"fmt"
)

func HowMuchSnow(mountMap []int) (int, []int) {

	// return early to avoid edge cases
	if len(mountMap) == 1 {
		return 1, []int{1}
	}

	// setup some vars
	snow := 0
	change := true
	snowMap := make([]int, len(mountMap))
	passCount := 0

	// loop over mountain until no more snow can be added
	for change {
		change = false

		// loop over mountain and add snow
		for i, block := range mountMap {
			switch i {

			// first block
			case 0:
				if snowMap[i] == 1 {
					// if the first index already has 1 block of snow do nothing
					continue
				} else if (snowMap[i+1]+mountMap[i+1])-block >= 0 {
					// if height of the mountain + the height of the snow in the adjacent index is >= add 1 block of snow
					snowMap[i] = 1
					snow++
					change = true
				}

			//last block
			case len(mountMap) - 1:
				if snowMap[i] == 1 {
					// if the last index already has 1 block of snow do nothing
					continue
				} else if snowMap[i-1]+mountMap[i-1]-block >= 0 {
					// if height of the mountain + the height of the snow in the adjacent index is >= add 1 block of snow
					snowMap[i] = 1
					snow++
					change = true
				}

			// middle blocks
			default:
				// if height of the mountain + the height of the snow in the adjacent indexes on both sides is >= add 1 block of snow
				left := snowMap[i-1] + mountMap[i-1]
				curr := snowMap[i] + block
				right := snowMap[i+1] + mountMap[i+1]
				if left >= curr && right >= curr {
					snowMap[i] += 1
					snow++
					change = true
				}
			}
		}

		passCount++
	}

	fmt.Printf("PassCount:%d\n", passCount)
	return snow, snowMap
}
