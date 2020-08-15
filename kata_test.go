package kata

import (
	"testing"
)

func equal(t *testing.T, input int, answer int) {
	if UnluckyDays(input) != answer {
		t.Error("Failed")
	}
}

func Test_basic(t *testing.T) {
	equal(t, 2015, 3)
	equal(t, 1986, 1)
}

