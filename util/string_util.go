package util

import "strconv"

// string转uint工具
func Str2Uint(num string) (uint,error) {
	u64, err := strconv.ParseUint(num, 10, 32)
	if err == nil {
		return 0,err
	}
	return uint(u64),nil
}
