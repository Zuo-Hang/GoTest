package main

import "fmt"

/*
 * Created by zuo_h on 23:54 2021/1/28
 */
//流程控制
func ifTest() {
	age := 19
	if age > 35 {
		fmt.Println("人到中年")
	} else if age > 18 {
		fmt.Println("澳门赌场开业了！")
	} else {
		fmt.Println("改写作业了")
	}

	//作用域 局部变量 释放内存，不堆积
	if age1 := 19; age1 > 35 {
		fmt.Println("人到中年")
	} else if age1 > 18 {
		fmt.Println("澳门赌场开业了！")
	} else {
		fmt.Println("改写作业了")
	}
}

func forTest() {
	//go 中只有一种循环语句，没有while
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//变种1
	var j = 0
	for ; j < 10; j++ {
		fmt.Println(j)
	}
	//变种2
	var m = 0
	for m < 10 {
		fmt.Println(m)
		m++
	}
	//无限循环
	//	for{
	//		fmt.Println(k)
	//	}

	//for range
	s := "hello 航航"
	for k, v := range s {
		fmt.Printf("%d,%c\n", k, v)
	}

	//输出9*9
	for w := 1; w < 10; w++ {
		for z := 1; z < w+1; z++ {
			fmt.Printf("%d*%d=%d\t", z, w, z*w)
		}
		fmt.Println()
	}

	//跳出for循环
	for num := 0; num < 10; num++ {
		if num == 5 {
			fmt.Println(num)
			break
		}
	}

	//跳过for循环
	for num := 0; num < 10; num++ {
		if num == 5 {
			continue
		}
		fmt.Println(num)
	}

}

func switchTest() {
	//局部变量
	switch i := 2; i {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小指")
	default:
		fmt.Println("无效数字")
	}

	i := 2
	switch i {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小指")
	default:
		fmt.Println("无效数字")
	}

	//变种 加判断
	var age int = 20
	switch {
	case age < 18:
		fmt.Println("好好学习")
	case 18 < age && age < 60:
		fmt.Println("好好工作")
	default:
		fmt.Println("好好活着")
	}

	//可以向前兼容 fallthrough 实现向下传递  但是不建议使用

}

//尽量少用 标签多的话 程序会异常晦涩难懂
func gotoTest() {

	//常规跳出多层循环
	b := false
	for w := 1; w < 10; w++ {
		for z := 1; z < w+1; z++ {
			fmt.Printf("%d*%d=%d\t", z, w, z*w)
			if w == 4 && z == 4 {
				b = true
				break
			}
		}
		fmt.Println()
		if b {
			break
		}
	}
	fmt.Println("常规做法")
	//goto
	for w := 1; w < 10; w++ {
		for z := 1; z < w+1; z++ {
			fmt.Printf("%d*%d=%d\t", z, w, z*w)
			if w == 4 && z == 4 {
				goto XX
			}
		}
		fmt.Println()
	}
XX: //标签
	fmt.Println("goto做法")
}

func main() {
	//ifTest()
	//forTest()
	//switchTest()
	gotoTest()
	//var i=1
	//j:=&i
	//fmt.Printf("j的类型是%T",j)
}
