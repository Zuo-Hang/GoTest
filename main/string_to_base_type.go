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
	//strconv.ParseBool()返回两个返回值，但是只想接收一个的时候，可以用下划线代替不想接收的返回值
	b, _ = strconv.ParseBool(str1)
	fmt.Printf("b'type is %T. b=%v\n", b, b)

	//string to int
	var str2 = "22"
	parseInt, _ := strconv.ParseInt(str2, 10, 64)
	fmt.Printf("parseInt'type is %T. b=%v\n", parseInt, parseInt)

	//string to float
	var str3 = "3.14"
	float, _ := strconv.ParseFloat(str3, 64)
	fmt.Printf("float'type is %T. float=%v", float, float)

}
