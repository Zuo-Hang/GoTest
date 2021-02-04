package main

import "fmt"

/*
 * Created by zuo_h on 0:04 2021/2/5
 */

func mapTest() {
	//无序的<k,v>结构
	//需要注意：1.要初始化   2.最好初始化的时候预估容量，避免动态扩容
	var m1 map[string]int
	fmt.Println(m1 == nil)
	m1 = make(map[string]int, 10) //类型 容量
	m1["理想"] = 18
	m1["jiwuming"] = 35
	fmt.Println(m1)

	//判断map中是否存在
	v, ok := m1["zhangSan"] //值，是否存在
	//不判断的话，可能出现空指针异常
	if ok {
		println(v)
	} else {
		println("不存在张三")
	}

	//遍历
	for k, v := range m1 {
		println(k, v)
	}
	//之遍历K
	for k := range m1 {
		println(k)
	}
	//只遍历V
	for _, v := range m1 {
		println(v)
	}

	//删除某个元素,若删除不存在的元素不出现异常（不操作）
	delete(m1, "jiwuming")
	fmt.Println(m1)

}

func main() {
	mapTest()
}
