package leetcode

import "testing"

func Test849(t *testing.T) {
	t.Log(maxDistToClosest2([]int{1, 0, 0, 1}))
}

func maxDistToClosest(seats []int) int {
	index, max, result := 0, 0, 0

	for i, v := range seats {
		if 1 == v {

			if i-index > max {
				max = i - index
			}
			index = i
		}
	}

	if len(seats)-1-index > max {
		max = len(seats) - 1 - index
	}

	if max%2 == 0 {
		result = max / 2
	} else {
		result = max/2 + 1
	}

	if index+max == len(seats)-1 || index == 0 {
		result = max
	}

	return result
}

func maxDistToClosest2(seats []int) int {
	res := 0
	l := 0
	for l < len(seats) && seats[l] == 0 {
		l++
	}
	res = max(res, l)
	for l < len(seats) {
		r := l + 1
		for r < len(seats) && seats[r] == 0 {
			r++
		}
		if r == len(seats) {
			res = max(res, r-l-1)
		} else {
			res = max(res, (r-l)/2)
		}
		l = r
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
