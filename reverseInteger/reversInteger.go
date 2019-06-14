package reverseInteger

import (
	"math"
	"strconv"
)

func reverseStr(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}

func reverse(x int) int {
	var negative bool
	if x < 0 {
		negative = true
		x = 0 - x
	}
	str := strconv.Itoa(x)
	r := reverseStr(str)
	res, _ := strconv.ParseInt(r, 10, 64)
	if negative {
		res = - res
	}
	if res < math.MinInt32 || res > math.MaxInt32{
		return 0
	} else {
		return int(res)
	}
}
