package mountain

import (
	"testing"
)

func TestGetMountainString(t *testing.T) {
	mountain := []int{1, 2, 3, 4, 3, 2, 1}
	expected := "   M   \n  MMM  \n MMMMM \nMMMMMMM\n"
	found := GetMountainString(mountain)
	if expected != found {
		t.Errorf("GetMountainString Failed, was expecting:\n%s got:\n%s\n", expected, found)
	}
}
