package leetcode

import "testing"

func Test2682(t *testing.T) {
	t.Log(circularGameLosers(5, 3))
}

// circularGameLosers @description: leetCodeUrl:https://leetcode.cn/problems/find-the-losers-of-the-circular-game/
// @parameter nums
// @parameter k
// @return []int
func circularGameLosers(n int, k int) []int {
	tempMap, i, j, result := map[int]int{}, 1, 1, make([]int, 0)

	for 0 == tempMap[j%n] {
		if j <= n {
			tempMap[j] = 1
		} else {
			tempMap[j%n] = 1
		}

		j = j + i*k
		i++
	}

	for t := 1; t <= n; t++ {
		if 0 == tempMap[t] {
			result = append(result, t)
		}
	}

	return result
}

func circularGameLosers2(n, k int) []int {
	visit := make([]bool, n)
	j := 0
	for i := k; !visit[j]; i += k {
		visit[j] = true
		j = (j + i) % n
	}
	ans := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if !visit[i] {
			ans = append(ans, i+1)
		}
	}
	return ans
}
