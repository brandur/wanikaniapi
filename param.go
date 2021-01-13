package wanikaniapi

import (
	"strconv"
	"time"
)

func Bool(b bool) *bool {
	return &b
}

func Int(i int) *int {
	return &i
}

func Time(t time.Time) *time.Time {
	return &t
}

func joinIDs(ids []ID, separator string) string {
	var s string

	for i, n := range ids {
		if i != 0 {
			s += ","
		}

		s += strconv.FormatInt(int64(n), 10)
	}

	return s
}

func joinInts(ints []int, separator string) string {
	var s string

	for i, n := range ints {
		if i != 0 {
			s += ","
		}

		s += strconv.Itoa(n)
	}

	return s
}

func joinStrings(strs []string, separator string) string {
	var s string

	for i, str := range strs {
		if i != 0 {
			str += ","
		}

		s += str
	}

	return s
}
