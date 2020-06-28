package kata

import (
	"testing"
)

func equal(t *testing.T, input int, answer int) {
	if NextHigher(input) != answer {
		t.Error("Failed")
	}
}

func Test_basic(t *testing.T) {
	equal(t, 129, 130)
	equal(t, 1, 2)
	equal(t, 323423, 323439)
	equal(t, 128, 256)
	equal(t, 1, 2)
	equal(t, 1022, 1279)
	equal(t, 127, 191)
	equal(t, 1253343, 1253359)
}

