package leetcode

import (
	"container/heap"
	"testing"
)

func Test2208(t *testing.T) {
	t.Log(halveArray([]int{5, 19, 8, 1}))
}

type PriorityQueue2 []float64

func (pq PriorityQueue2) Len() int {
	return len(pq)
}

func (pq PriorityQueue2) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq PriorityQueue2) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue2) Push(x any) {
	*pq = append(*pq, x.(float64))
}

func (pq *PriorityQueue2) Pop() any {
	old, n := *pq, len(*pq)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func halveArray(nums []int) int {
	pq := &PriorityQueue2{}
	sum, sum2 := 0.0, 0.0
	for _, x := range nums {
		heap.Push(pq, float64(x))
		sum += float64(x)
	}
	res := 0
	for sum2 < sum/2 {
		x := heap.Pop(pq).(float64)
		sum2 += x / 2
		heap.Push(pq, x/2)
		res++
	}
	return res
}

func halveArray2(nums []int) int {
	sum, rest, max, count := 0, float32(0), float32(0), 0
	numss := make([]float32, 0)

	for _, num := range nums {
		sum += num
		numss = append(numss, float32(num))
	}
	rest = float32(sum)

	for rest > float32(sum)/2 {
		index := 0
		for i, num := range numss {
			if max < num {
				max = num
				index = i
			}
		}
		rest -= float32(max) / 2
		numss[index] = float32(max) / 2
		count++
		max = 0
	}

	return count
}
