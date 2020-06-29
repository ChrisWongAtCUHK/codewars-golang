package kata

import (
	"math/rand"
	"testing"
	"time"
)


func sol(n int) int {
    o := n & -n
    v := n + o
    return v | ((n ^ v) / o >> 2)
}

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

func Test_random(t * testing.T) {
	rand.Seed(time.Now().UnixNano())
        m := (1 << 30) - 1
        for i := 0; i < 100; i++ {
            n := rand.Intn(m) + 1
			equal(t, n, sol(n))
		}
}

