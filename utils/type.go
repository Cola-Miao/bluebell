package utils

import "strconv"

func ToAnySlice[T any](s []T) []any {
	l := len(s)
	res := make([]any, l)
	for i := 0; i < l; i++ {
		res[i] = s[i]
	}
	return res
}

func AtoInt64(s []string) ([]int64, error) {
	l := len(s)
	res := make([]int64, l)
	for i := 0; i < l; i++ {
		num, err := strconv.Atoi(s[i])
		if err != nil {
			return nil, err
		}
		res[i] = int64(num)
	}
	return res, nil
}
