package multipass

import (
	"testing"
)

func TestHowMuchSnow(t *testing.T) {
	CheckMountain(t, []int{1}, 1)
	CheckMountain(t, []int{1, 1}, 2)
	CheckMountain(t, []int{5, 1}, 1)
	CheckMountain(t, []int{1, 2, 1}, 3)
	CheckMountain(t, []int{2, 1, 2}, 5)
	CheckMountain(t, []int{1, 2, 3, 2, 1}, 5)
	CheckMountain(t, []int{2, 1, 1, 2}, 8)
	CheckMountain(t, []int{1, 2, 1, 2, 1}, 7)
	CheckMountain(t, []int{2, 1, 2, 1, 2}, 11)
	CheckMountain(t, []int{1, 3, 3, 3, 1}, 3)
	CheckMountain(t, []int{1, 4, 4, 4, 2}, 3)
	CheckMountain(t, []int{96, 54, 86, 35, 10, 75, 7}, 142)
	CheckMountain(t, []int{7, 75, 10, 35, 86, 54, 96}, 142)
}

func CheckMountain(t *testing.T, mount []int, expected int) {
	found, _ := HowMuchSnow(mount)
	if found != expected {
		t.Errorf("Expected: %d Found: %d", expected, found)
	}
}
