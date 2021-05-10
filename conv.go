package dateparse

import (
	"strconv"
	"strings"
)

func forceInt(s string) int { return int(forceInt64(s)) }

func forceInt64(s string) int64 {
	s = strings.TrimSpace(s)
	s = strings.TrimLeft(s, "0")
	if s == "" {
		return 0
	}
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return val
}

func normalizeStrings(ss []string) []string {
	if len(ss) <= 1 {
		return ss
	}
	res := make([]string, 0, len(ss))
	for _, bit := range ss {
		bit = strings.TrimSpace(bit)
		if bit != "" {
			res = append(res, bit)
		}
	}
	return res
}
