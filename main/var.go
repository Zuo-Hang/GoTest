package main

import fmt "fmt"

/*
 * Created by zuo_h on 0:26 2021/1/19
 */

//全局变量

var m1 = 100
var m2 = 200
var m3 = 300

var m4 = "wangWu"

//一次性声明

var (
	h1 = 21
	h2 = "wangMaZi"
)

func main() {
	fmt.Println(m1, m2, m3, m4)
	fmt.Println(h1, h2)
	//局部变量
	var i int
	i = 1
	fmt.Println(i)

	var j int
	fmt.Println(j)

	var k = 1.1
	fmt.Println("k=", k)

	//类型推导
	w := "tom"
	fmt.Println("w=", w)

	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

	var y1, y2, y3 = 20, "zhangSan", 1.74
	fmt.Println(y1, y2, y3)

	z1, z2, z3 := 20, "liSi", 1.80
	fmt.Println(z1, z2, z3)

}
