package atoi

import (
	"math"
	"strings"
)

func atoi(str string) int {
	num := strings.Fields(str)
	if len(num) == 0 {
		return 0
	}
	res := parseInt(num[0])
	if res < math.MinInt32 {
		return math.MinInt32
	}
	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	return int(res)
}

func parseInt(s string) int64 {
	if len(s) == 0 {
		return 0
	}
	neg := false
	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		neg = true
		s = s[1:]
	}
	var res int64
	count := 0
	for _, n := range s {
		if (n >= '0') && (n <= '9') {
			res = res*10 + int64(n-48)
			if res != 0 {
				count++
			}
		} else {
			break
		}
		if count > 10 {
			break
		}
	}
	if neg {
		return -res
	}
	return res
}
