package main

import "fmt"

/*
 * Created by zuo_h on 0:25 2021/1/30
 */
//切片
//切片底层是指向了一个数组
//切片的长度即长度  容量是从切片的第一个元素到底层数组的最后一个元素的数量
//切片可以扩容
//切片可以再切片
//底层数组改变了，切片的值也改变(引用类型)
func slice() {
	//定义
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)
	//nil (没有开辟内存空间) 是 null 的意思
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	//初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"北京", "上海", "广州", "深圳"}
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	//长度和容量
	fmt.Println(len(s1), cap(s1))

	//由数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s3 := a1[0:4] //索引  左包含，右不包含
	fmt.Println(s3)
	s4 := a1[:4]
	s5 := a1[4:] //a1[4:len(a1)]
	s6 := a1[:]
	//容量是指
	fmt.Println(s4, len(s4), cap(s4))
	fmt.Println(s5, len(s5), cap(s5))
	fmt.Println(s6, len(s6), cap(s6))

}

func makeTest() {
	//切片的本质就是一个框，框住了一块连续的内存。属于引用类型，真正的值都存在数组里面。
	//切片不能进行比较，只能和nil比较，一个nil的切片是没有底层数组的，len、和cap都是0。但是我们不能说一个冷、和cap都是0的切片一定是nil
	//判断一个切片是空，使用len()==0
	//注意数组只要定义了就会开辟内存空间，与Java不同。可以说数组不是引用类型
	//make函数用来创造一个切片，
	s1 := make([]int, 5, 10) //类型、长度、容量
	fmt.Printf("s1=%v,s1.type=%T,s1.len()=%d,s1.cap()=%d", s1, s1, len(s1), cap(s1))

}

func appendTest() {
	//append() 追加

}

func main() {
	//slice()
	makeTest()
}
