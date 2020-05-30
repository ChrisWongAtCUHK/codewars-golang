package kata

import (
	"reflect"
	"testing"
)

func Test_1(t *testing.T) {
	result := SplitAndAdd([]int{1, 2, 3, 4, 5}, 2)
	if reflect.DeepEqual(result, []int{5, 10}) == false {
		t.Error("Failed")
	}
}

