package leetcode

import "testing"

func Test2337(t *testing.T) {
	t.Log(canChange("_R", "R_"))
}

func canChange(start string, target string) bool {
	i, j, n := 0, 0, len(start)
	for i < n && j < n {
		for i < n && start[i] == '_' {
			i++
		}
		for j < n && target[j] == '_' {
			j++
		}
		if i < n && j < n {
			if start[i] != target[j] {
				return false
			}
			c := start[i]
			if c == 'L' && i < j || c == 'R' && i > j {
				return false
			}
			i++
			j++
		}
	}
	for i < n {
		if start[i] != '_' {
			return false
		}
		i++
	}
	for j < n {
		if target[j] != '_' {
			return false
		}
		j++
	}
	return true
}
