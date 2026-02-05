package kyu8

import "strconv"

func SumMix(arr []any) int {
	sum := 0
	for _, item := range arr {
		switch v := item.(type) {
		case int:
			sum += v
		case string:
			if num, err := strconv.Atoi(v); err == nil {
				sum += num
			}
		}
	}
	return sum
}

//https://www.codewars.com/kata/57eaeb9578748ff92a000009/solutions/go
