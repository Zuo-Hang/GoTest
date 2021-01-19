package main

/*
 * Created by zuo_h on 23:08 2021/1/19
 */
import (
	"fmt"
	"unsafe"
)

func main() {
	//Base Data Type

	//有符号整数 int8(-128~127) int16 int32 int64 取值范围
	//int 和操作系统有关 32位系统4字节 64位8字节 (默认)
	var i1 int8 = 127
	var i3 int64 = 8
	fmt.Println(i1)
	//无符号整数 uint8(0~255) uint16 uint32 uint64
	var i2 uint8
	fmt.Println(i2)

	var b byte = 255
	fmt.Println(b)

	//查看数据类型
	fmt.Printf("b 的数据类型是 %T \n", b)
	//查看变量占用字节大小
	fmt.Printf("i3 的数据类型是 %T 占用字节大小 %d", i3, unsafe.Sizeof(i3))

	//分配类型的时候需要保小不保大原则 占用较小空间

	//bit

	//浮点类型float32 4字节 float64 8字节  可能有精度损失  浮点数默认分配的是64位格式

	var price float32 = 89.6
	fmt.Println(price)
	var price2 float32 = -89.6
	fmt.Println(price, price2)

	var accuracy1 float32 = 89.690008901
	fmt.Println(price)
	var accuracy2 float64 = 89.690008901
	fmt.Println(accuracy1, accuracy2)

	//科学计数法
	num1 := 2.569e2
	num2 := 2.569e2
	num3 := 2.569e-2
	fmt.Println(num1, num2, num3)

	//字符类型 没有char类型 存储单个字符使用byte（由于是utf-8所以不可保存中文） 字符串由字符组成，而不是字节
	//直接输出的话 输出的是其ASCII 所以其取值范围也限定在其中
	var c1 byte = 'a'
	var c2 byte = '0'
	fmt.Println(c1, c2)
	//按字符输出需要格式化
	fmt.Printf("%c %c", c1, c2)
	//汉字的存储
	var c3 int = '北'
	fmt.Printf("%c %d", c3, c3)

}
