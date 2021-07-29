/*
 @Author : dongyang
 @Time   : 2021/7/29 10
 @Desc   :
*/
package common

import "strings"

// 对错误包装，进行可视化展示
func ErrPos(s string, pos int) string {
	r := strings.Repeat("-", len(s)) + "\n"
	s += "\n"
	for i := 0; i < pos; i++ {
		s += " "
	}
	s += "^\n"
	return r + s + r
}
