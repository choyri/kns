package support

import (
	"strconv"
)

func Str2Uint(value string) uint64 {
	ret, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		ret = 0
	}

	return ret
}
