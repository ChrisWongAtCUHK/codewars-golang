package kata

// AvramenkoArtem
import (
	"strconv"
	"strings"
)

func NextHigher(n int) int {
	for i := n + 1; i > 0; i++ {
		if strings.Count(strconv.FormatInt(int64(n), 2), "1") == strings.Count(strconv.FormatInt(int64(i), 2), "1") {
			return i
		}
	}
	return 0
}
