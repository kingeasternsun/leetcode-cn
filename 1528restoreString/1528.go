/*
 * @Description: 1528
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-06 10:24:45
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-06 11:07:35
 * @FilePath: \1528restoreString\1528.go
 */
package leetcode

func restoreString(s string, indices []int) string {
	in := []byte(s)

	for i := 0; i < len(in); i++ {

		//位置没变就不用动
		if indices[i] == i {
			continue
		}
		ch := in[i]
		nextID := indices[i]
		//如果下面的 for 循环用 indices[nextID] != nextID作为终止条件，这个就必须,不然就会出现重复处理的情况
		indices[i] = i
		for indices[nextID] != nextID { //判断条件就是 indices[nextID] != nextID
			//放入到 in的 nextIDn 位置上，并且把 in 在 nextIDn 位置上原来的value保存到ch
			in[nextID], ch = ch, in[nextID]
			// curID 赋给 indices[nextID] 标识已经更换过不用再动了，同时把 更新curID 为 indices[nextID]
			indices[nextID], nextID = nextID, indices[nextID]
		}

		//当前i的位置要用最新的ch来填充
		in[i] = ch
		indices[i] = i
	}

	return string(in)
}
