package main

import (
	"fmt"
	"strconv"
)

/*
 * Created by zuo_h on 23:18 2021/1/27
 */

//运算符 六种 ：算术运算 赋值运算 比较运算（关系运算） 逻辑运算 位运算 其他运算（& *） 在go中没有三元运算符
func arithmetic() {
	//算术运算符：+ - * / % ++ -- +（字符串拼接）
	fmt.Println(10 / 3)
	fmt.Println(10 % 4)
	fmt.Println(10 % -3)
	var i int = 10
	//在Go中 ++ -- 只能当作一个独立的语言来使用 不能 k:=i++ 这样组合使用
	//Go中的 ++ -- 只能跟在变量后 不存在 ++1 这种前置自增
	i++
	fmt.Println(i)
}

func relationship() {
	//关系运算符
	// ==  !=  >  <  >=  <=
	b := 7 > 8
	fmt.Println(b)
}

func logic() {
	//逻辑运算符：将多个关系运算链接
	//&& || !  短路
	var age = 30
	if age > 20 && age < 40 {
		fmt.Println(strconv.Itoa(age) + " between 20 and 40")
	}
	if age > 20 || age < 10 {
		fmt.Println(strconv.Itoa(age) + " big than 20 or small than 10")
	}

}

func bitwiseOperator() {
	//位运算
	//5:101 2:10
	fmt.Println(5 & 2)  //000
	fmt.Println(5 | 2)  //111
	fmt.Println(5 ^ 2)  //异或 111
	fmt.Println(5 << 1) //左移 1010
	fmt.Println(5 >> 1) //右移 10

	var i = int8(1)
	fmt.Println(i << 10) //溢出 使用的时候注意

}

func assignment() {
	//赋值运算符
	// = 符合运算符 += -= *= /= %= <<= >>= &= |= ^=
	var a = 10
	var b = 20
	fmt.Printf("a=%v,b=%v\n", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a=%v,b=%v\n", a, b)

}

func main() {
	//arithmetic()
	//relationship()
	//logic()
	//assignment()
	bitwiseOperator()
}
