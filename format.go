package money

import (
	"fmt"
	"strings"
)

//Format 格式化金额.
func Format(money float64) string {
	m := fmt.Sprintf("%.2f", money)
	arr := strings.Split(m, ".")
	sarr := strings.Split(arr[0], "")
	var str string
	var length int
	for i := len(sarr); i > 0; i-- {
		num := i - 1
		length++
		str = fmt.Sprintf("%s%s", sarr[num], str)
		if length%3 == 0 {
			str = fmt.Sprintf(",%s", str)
		}
	}
	return fmt.Sprintf("%s.%s", strings.TrimLeft(str, ","), arr[1])
}
