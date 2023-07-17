package leetcode

import (
	"strconv"
	"testing"
)

func Test415(t *testing.T) {
	t.Log(addStrings("11", "123"))
}

// addStrings @description: leetCodeUrl:https://leetcode.cn/problems/add-strings/
// @parameter num1
// @parameter num2
// @return string
func addStrings(num1 string, num2 string) string {
	result := ""
	if len(num1) > len(num2) {
		num2 = addZero(len(num1)-len(num2)) + num2
	} else {
		num1 = addZero(len(num2)-len(num1)) + num1
	}

	list, is := make([]int, 0), false
	for i := len(num1); i > 0; i-- {
		num := int(num1[i-1]-'0') + int(num2[i-1]-'0')
		if is {
			num++
			is = false
		}
		list = append(list, num%10)
		if num >= 10 {
			is = true
		}
	}

	if is {
		list = append(list, 1)
	}

	for i := len(list) - 1; i >= 0; i-- {
		result += strconv.Itoa(list[i])
	}
	return result
}

func addZero(num int) (result string) {
	for i := 0; i < num; i++ {
		result += "0"
	}
	return
}
