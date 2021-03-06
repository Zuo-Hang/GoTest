package main

/*
 * Created by zuo_h on 23:30 2021/1/25
 * 指针示例
 */

import (
	"GoTest/src/private"
	"fmt"
)

func point() {
	//与C不同的是 go不可以进行指针运算
	// &取地址符  *int 指针类型
	i := 23
	fmt.Println("i 的地址是", &i)
	var p *int = &i
	fmt.Printf("p's type is %T. p=%v\n", p, p)

	fmt.Println("p 的地址是", &p)
	fmt.Println("p 的值是", p)
	fmt.Println("p 指向的值是", *p)

	//利用指针修改值
	*p = 64
	fmt.Println("i 的值是", i)

	//go中的值类型和引用类型
	//值类型：基本数据类型 int系列、float系列、bool、string、数组、结构体struct
	//引用类型：指针、slice切片、map、管道chan、interface等
	//值类型直接存储，一般分配在栈中。引用类型通常在堆区分配（逃逸分析除外）

	//标识符
	//英文、数字、下划线
	//数字不可开头
	//区分大小写
	//单独的下划线 _ 代表空，可以代表任何一个标识符，但是它的值会被忽略，故它只能做占位符不能做标识符
	//注意25个关键字

	//包名尽量和文件夹名相同、不要和标准库重名（可以不同，但是标准是相同的）
	//变量、方法名、常量名——>驼峰法
	//** 首字母大写为共有的（public意义）、首字母小写的是私有的（private意义）
	//导包 ——> 在创建项目的时候，要注意go的路径结构、配置GOPATH。要不然会出现无法导包的问题。
	//|--workspace
	//----|src
	//------|project
	//----|pkg
	//----|bin
	name := private.Name
	//private.name  这个是不能导入的
	fmt.Println("private 包下的 Name 是", name)
}

func newAndMake() {
	//这是不对的，虽然定义了一个指针变量，但是并没有赋初始值（没有开辟内存空间），为默认值nil。当下面通过指针赋值的时候，
	//要通过地址获取内存时就会出现空指针异常
	//var a *int
	//*a=100
	//fmt.Println(*a)

	//new函数开辟空间（一般用于基本类型），返回的是对应类型的指针  基本很少用到
	var a = new(int)
	*a = 100
	fmt.Println(*a)

	//make也用于开辟内存空间，但是区别与new  只用于slice、map、chan的内存创建，返回的不是指针，而是这几个数据类型本身

}

func main() {
	point()
	newAndMake()
}
