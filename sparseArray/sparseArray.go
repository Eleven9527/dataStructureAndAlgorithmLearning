package sparseArray

import "fmt"

/*
0 0 0 0 0
0 1 0 0 0
0 0 2 0 0
0 0 0 0 0
0 0 0 0 0
*/
/*
row col val
5	5	2
1	1	1
2	2	2
*/
func SparseArrayTest() {
	fmt.Println("----------------origin array-----------------------")
	oa := initOriginArray()
	for _, v := range oa {
		fmt.Println(v)
	}

	fmt.Println("---------------to sparse array---------------------")
	sa := initSparseArray(oa)
	for _, v := range sa {
		fmt.Println(v)
	}

	fmt.Println("---------------to origin array---------------------")
	oa2 := recoverToOriginArray(sa)
	for _, v := range oa2 {
		fmt.Println(v)
	}

	fmt.Println(oa == oa2)
}

//初始化二维数组
func initOriginArray() [5][5]int {
	var oa [5][5]int
	oa[1][1] = 1
	oa[2][2] = 2

	return oa
}

//初始化稀疏数组
func initSparseArray(ca [5][5]int) [][3]int {
	sa := make([][3]int, 0)
	sa = append(sa, [3]int{})

	var count int //非0元素的个数

	for i, _ := range ca {
		for j, v := range ca[i] {
			if v != 0 {
				count++
				temp := [3]int{i, j, v}
				sa = append(sa, temp)
			}
		}

		overView := [3]int{len(ca), len(ca[0]), count}
		sa[0] = overView
	}

	return sa
}

//稀疏数组->二维数组
func recoverToOriginArray(sa [][3]int) (ret [5][5]int) {
	for k, v := range sa {
		if k == 0 {
			continue
		}

		ret[v[0]][v[1]] = v[2]
	}

	return
}
