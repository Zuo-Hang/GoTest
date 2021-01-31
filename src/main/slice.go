package main

import (
	"fmt"
	"sort"
)

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
	fmt.Printf("s1=%v,s1.type=%T,s1.len()=%d,s1.cap()=%d\n", s1, s1, len(s1), cap(s1))

}

func appendTest() {
	//append() 追加 扩容 ***会修改底层数组
	s1 := []string{"北京", "上海", "广州"}
	//s1[3]="西安" 错误的写法，会导致编译错误：索引越界(ut of range [4] with length 4 )。这样做无法进行扩容
	fmt.Printf("s1=%v,s1.type=%T,s1.len()=%d,s1.cap()=%d\n", s1, s1, len(s1), cap(s1))
	//扩容函数必须用被扩容的变量接收,底层会重新创建一个数组
	s1 = append(s1, "西安") //追加的内容
	fmt.Printf("s1=%v,s1.type=%T,s1.len()=%d,s1.cap()=%d\n", s1, s1, len(s1), cap(s1))
	//扩容机制：
	//	最早是直接2倍，后面进行了优化
	//	首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。
	//	否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
	//	否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
	//	如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。
	//	需要注意的是，切片扩容还会根据切片中元素的类型不同而做不同的处理，比如int和string类型的处理方式就不一样。
	s2 := []string{"苏州", "武汉", "成都"}
	s1 = append(s1, s2...) // ... 表示将这个slice拆开
	fmt.Printf("s1=%v,s1.type=%T,s1.len()=%d,s1.cap()=%d\n", s1, s1, len(s1), cap(s1))
}

func copyTest() {
	a1 := []int{1, 9, 3}
	a2 := a1
	var a3 = make([]int, 3, 3)
	//拷贝要求目标具有容量和长度
	copy(a3, a1)
	fmt.Println(a1, a2, a3)
	a1[2] = 6
	fmt.Println(a1, a2, a3)
	//切片中没有删除元素的方法，只能将原切片再切片成两个（不包含此元素的），进行拼接。(类似前移一个元素)
	//移除第二个元素（索引为1）
	a1 = append(a1[:1], a2[2:]...)
	fmt.Println(a1)
	sort.Ints(a3) //对切片排序
	fmt.Println(a3)

}

func main() {
	//slice()
	//makeTest()
	//appendTest()
	copyTest()
}
