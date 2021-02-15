/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-11 23:52:51
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-12 00:07:49
 * @FilePath: /703KthLargest/703_test.go
 */
package leetcode

import "testing"

func TestK(t *testing.T) {
	obj := Constructor(3, []int{4, 5, 8, 2})
	param_1 := obj.Add(3)
	if param_1 != 4 {
		t.Error("ee")
	}
	param_2 := obj.Add(5)
	if param_2 != 5 {
		t.Error("ee5")
	}
}

func TestK1(t *testing.T) {
	obj := Constructor(1, []int{})
	param_1 := obj.Add(3)
	if param_1 != 3 {
		t.Error("ee")
	}
	param_2 := obj.Add(5)
	if param_2 != 5 {
		t.Error("ee5")
	}
}

func TestK2(t *testing.T) {
	obj := Constructor(2, []int{0})
	param_1 := obj.Add(-1)
	if param_1 != -1 {
		t.Error("ee")
	}
	param_2 := obj.Add(1)
	if param_2 != 0 {
		t.Error("ee5")
	}
}
