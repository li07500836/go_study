package valuechangepackage

import (
	"strconv"
)

func IntToString(i int) string  {
	var reString string = strconv.Itoa(int(i))
	return reString
}