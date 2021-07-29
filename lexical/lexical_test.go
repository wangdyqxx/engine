/*
 @Author : dongyang
 @Time   : 2021/7/28 20
 @Desc   :
*/
package lexical

import "testing"

func TestMath_Parse(t *testing.T) {
	str := "1. 2 .3"
	ts,err:= Parse(str)
	t.Log(ts,err)
	for _, token := range ts {
		t.Log(token)
	}
}
