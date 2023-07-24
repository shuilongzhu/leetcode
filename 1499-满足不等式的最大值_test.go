package leetcode

import (
	"container/heap"
	"testing"
)

func Test1499(t *testing.T) {
	t.Log(findMaxValueOfEquation([][]int{{1, 3}, {2, 0}, {5, 10}, {6, -10}}, 1))
}

type PriorityQueue [][]int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i][0] != pq[j][0] {
		return pq[i][0] < pq[j][0]
	}
	return pq[i][1] < pq[j][1]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.([]int))
}

func (pq *PriorityQueue) Pop() any {
	n, old := len(*pq), *pq
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func (pq PriorityQueue) Top() []int {
	return pq[0]
}

// findMaxValueOfEquation @description:leetCodeUrl:https://leetcode.cn/problems/max-value-of-equation/submissions/
// @parameter points
// @parameter k
// @return int
func findMaxValueOfEquation(points [][]int, k int) int {
	res := -0x3f3f3f3f
	pq := &PriorityQueue{}
	for _, point := range points {
		x, y := point[0], point[1]
		for pq.Len() > 0 && x-pq.Top()[1] > k {
			heap.Pop(pq)
		}
		if pq.Len() > 0 {
			res = Max(res, x+y-pq.Top()[0])
		}
		heap.Push(pq, []int{x - y, x})
	}
	return res
}

func findMaxValueOfEquation2(points [][]int, k int) int {
	result, state := 0, 0

	for i, _ := range points {

		for j := i + 1; j < len(points); j++ {

			if points[j][0]-points[i][0] > k {
				break
			}
			if 0 == state {
				result = points[i][1] + points[j][1] + points[j][0] - points[i][0]
			}
			state = 1
			if result < points[i][1]+points[j][1]+points[j][0]-points[i][0] {
				result = points[i][1] + points[j][1] + points[j][0] - points[i][0]
			}
		}
	}

	return result
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
