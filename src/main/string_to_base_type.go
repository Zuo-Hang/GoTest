package main

/*
 * Created by zuo_h on 23:02 2021/1/25
 */

import (
	"fmt"
	"strconv"
)

func main() {
	//string to bool
	var str1 string = "true"
	fmt.Printf("str1'type is %T\n", str1)
	var b bool
	//strconv.ParseBool()返回两个返回值num,err.但是只想接收一个的时候，可以用下划线代替不想接收的返回值
	b, _ = strconv.ParseBool(str1)
	fmt.Printf("b'type is %T. b=%v\n", b, b)

	//string to int
	var str2 = "22"
	parseInt, _ := strconv.ParseInt(str2, 10, 64)
	fmt.Printf("parseInt'type is %T. parseInt=%v\n", parseInt, parseInt)

	//string to float
	var str3 = "3.14"
	float, _ := strconv.ParseFloat(str3, 64)
	fmt.Printf("float'type is %T. float=%v\n", float, float)

	//go 在数据类型转换的时候，要是不是一个有效数字，则会将转换后的数字类型的默认值赋值给结果
	parseInt1, _ := strconv.ParseInt(str1, 10, 64)
	fmt.Printf("parseInt1'type is %T. parseInt1=%v\n", parseInt1, parseInt1)

	//字符串更改    ----切片  切片里面保存的是字符： byte rune  (其实都是int32)
	var s1 = "白萝卜"
	s2 := []rune(s1)
	s2[0] = '红'
	fmt.Println(string(s2))
	c1 := "白"
	c2 := '白'
	c3 := "H"
	c4 := 'H'
	fmt.Printf("c1=%T\n", c1)
	fmt.Printf("c2=%T\n", c2)
	fmt.Printf("c3=%T\n", c3)
	fmt.Printf("c4=%T\n", c4)

}
