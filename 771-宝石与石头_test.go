package leetcode

import (
	"fmt"
	"testing"
)

func Test771(t *testing.T) {
	t.Log(numJewelsInStones("aA", "aAAbbbb"))
}

func numJewelsInStones(jewels string, stones string) int {
	jewelsCount := 0
	jewelsSet := map[byte]bool{}
	t := jewels[0]
	fmt.Println(t)
	for i := range jewels {
		jewelsSet[jewels[i]] = true
	}
	for i := range stones {
		if jewelsSet[stones[i]] {
			jewelsCount++
		}
	}
	return jewelsCount
}
