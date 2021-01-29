package main

import "fmt"

/*
 * Created by zuo_h on 23:45 2021/1/29
 */

//类似其他语言的列表，存放基本的数据类型。需要定义类型，
func array() {
	//数组的长度是数组类型的一部分，比如a1和a2无法比较.所以数组的使用比较少
	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("a1 是 %T类型\n", a1)
	fmt.Printf("a2 是 %T类型\n", a2)

	//数组的初始化
	fmt.Println(a1, a2) //默认数值为0值
	//初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	//初始化方式2  根据初始化的值自动判断容量
	a10 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a10)
	//初始化方式3
	a3 := [10]int{1, 2, 3}
	fmt.Println(a3)
	//初始化方式4 根据索引初始化
	a4 := [10]int{0: 1, 9: 9}
	fmt.Println(a4)

	//数组是值类型  支持==  ！= 操作
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)
}

func ergodic() {
	//根据索引遍历
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < 10; i++ {
		fmt.Println(a1[i])
	}
	//for range
	for i, v := range a1 {
		fmt.Printf("%d,%d\n", i, v)
	}
}

func twoDimensionalArray() {
	//定义
	var a [2][3]int
	a = [2][3]int{
		[3]int{1, 2, 3},
		[3]int{4, 5, 6},
	}
	fmt.Println(a)
	//遍历
	for _, i := range a {
		for _, j := range i {
			fmt.Println(j)
		}
	}
}

func main() {
	array()
	ergodic()
	twoDimensionalArray()
}
