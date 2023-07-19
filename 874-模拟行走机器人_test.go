package leetcode

import (
	"strconv"
	"testing"
)

func Test874(t *testing.T) {
	t.Log(robotSim([]int{4, -1, 4, -2, 4}, [][]int{{2, 4}}))
}

// robotSim @description: leetCodeUrl:https://leetcode.cn/problems/walking-robot-simulation/
// @parameter commands
// @parameter obstacles
// @return int
func robotSim(commands []int, obstacles [][]int) int {
	lists := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	result, x, y, b, maps := 0, 0, 0, 1, make(map[string]int, 0)

	//障碍物数组转化为map,方便查找
	for _, obstacle := range obstacles {
		maps[strconv.Itoa(obstacle[0])+"="+strconv.Itoa(obstacle[1])] = 1
	}

	for _, c := range commands {
		if -1 == c { //右转90°
			b = (b + 1) % 4
			continue
		}
		if -2 == c { //左转90°
			b = (b + 3) % 4
			continue
		}
		for i := 0; i < c; i++ {
			if 1 == maps[strconv.Itoa(x+lists[b][0])+"="+strconv.Itoa(y+lists[b][1])] {
				break
			}
			x += lists[b][0]
			y += lists[b][1]
			result = max(result, x*x+y*y)
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
