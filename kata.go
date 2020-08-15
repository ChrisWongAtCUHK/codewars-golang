package kata

import (
	"time"
)

func UnluckyDays(year int) int {
	count := 0;
	t := time.Date(year, time.January, 13, 0, 0, 0, 0, time.UTC)
	for i := 0; i < 11; i++ {
		if int(t.AddDate(0, i, 0).Weekday()) == 5 {
			count = count + 1
		}
	}

	return count;
}

