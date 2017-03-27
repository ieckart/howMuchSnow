package mountain

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	MountChar = "█"
	SnowChar  = "❄"
	MaxHeight = 30 // real maxHeight == maxHeight + 1
	MaxWidth  = 99 // real maxWidth == maxWidth + 1
)

// Returns an int slice that represents the topographical map of a randomly generated mountain range.
func GetMountain() ([]int, int64) {
	return GetMountainS(time.Now().UTC().UnixNano())
}

// Returns an int slice that represents the topographical map of a randomly generated mountain range with a given seed.
func GetMountainS(seed int64) ([]int, int64) {
	rand.Seed(seed)
	return GetMountainSWH(seed, rand.Intn(MaxWidth)+1, rand.Intn(MaxHeight)+1)
}

// Returns an int array that represents the topographical map of a randomly generated mountain range with a given width
func GetMountainW(width int) ([]int, int64) {
	return GetMountainSWH(time.Now().UTC().UnixNano(), width, MaxHeight)
}

// Returns an int array that represents the topographical map of a randomly generated mountain range with a fixed width, max height, and given seed
func GetMountainSWH(seed int64, width, height int) ([]int, int64) {
	rand.Seed(seed)
	if width == 0 {
		width = 1
	}
	tMap := make([]int, width)
	for i, _ := range tMap {
		h := rand.Intn(height)
		if h == 0 {
			h = 1
		}
		tMap[i] = h
	}
	return tMap, seed
}

// returns a nicely formatted string representation of a mountain range
func GetMountainString(mountain []int) string {
	maxHeight := 0
	for _, val := range mountain {
		if val > maxHeight {
			maxHeight = val
		}
	}

	mountString := ""
	for h := maxHeight; h > 0; h-- {
		for _, val := range mountain {
			if val >= h {
				mountString = mountString + MountChar
			} else {
				mountString = mountString + " "
			}
		}
		mountString = mountString + "\n"
	}
	return mountString
}

// returns a nicely formatted string representation of a mountain range with snow on it
func GetMountainSnowString(mountain, snow []int) string {
	if len(mountain) != len(snow) {
		return "Mountain and snow arrays must be on the same size :*("
	}
	maxHeight := 0
	for i, val := range mountain {
		if val+snow[i] > maxHeight {
			maxHeight = val + snow[i]
		}
	}

	mountString := ""
	for h := maxHeight; h > 0; h-- {
		for i, val := range mountain {
			if val >= h {
				mountString = mountString + MountChar
			} else if val+snow[i] >= h {
				mountString = mountString + SnowChar
			} else {
				mountString = mountString + " "
			}
		}
		mountString = mountString + "\n"
	}
	return mountString
}

// returns an array with [x, y, z] format, not [x y z]
func FormatArr(arr []int) string {
	return strings.Replace(fmt.Sprintf("%v", arr), " ", ", ", -1)
}
