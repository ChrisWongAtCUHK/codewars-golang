package kata

import (
	"reflect"
	"testing"
)

func Test_1(t *testing.T) {
	result := NextHigher(128)
	if reflect.DeepEqual(result, 256) == false {
		t.Error("Failed")
	}
}

