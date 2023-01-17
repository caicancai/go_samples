package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sliceDefine()
}

// 数组定义、赋值、遍历
func arrayTest() {
	//声明
	var arr1 [5]string
	//赋值
	arr1[0] = "are"
	arr1[1] = "you"
	arr1[2] = "ok"
	//获取长度
	fmt.Println(len(arr1))
	//遍历-1
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("%d:%s \n", i, arr1[i])
	}
	//遍历-2
	for i, e := range arr1 {
		fmt.Printf("%d:%s \n", i, e)
	}
}

// 数组拷贝，默认的值拷贝，和引用拷贝方式
func arrayCopy() {
	arr1 := [3]int{1, 2, 3}
	//go数组默认是值拷贝
	arr2 := arr1 // arr2 是对 arr1 的值拷贝，对 arr2 的修改不会传递到 arr1
	arr2[0] = 233
	fmt.Println(arr1, arr2)

	//引用拷贝数组
	arr3 := &arr1 //数组arr3是arr1的指针引用，对arr3的修改会传递到arr1
	arr3[0] = 233
	fmt.Println(arr1, arr3) // [233 2 3] &[233 2 3]
}

// 数组常量，声明数组时进行赋值
func arrayConst() {
	var arr1 = [3]int{1, 2, 3}
	var arr2 = [10]int{1, 2, 3} // arr2 10 个元素，从第4个元素开始都是默认值0

	arr3 := [...]int{1, 2, 3}
	arr4 := []int{1, 2, 3}

	arr5 := [5]string{3: "are", 4: "you"} //指定元素及其索引值

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5, arr5[0] == "")
}

// 多维数组
func matrixTest() {
	var pixel [10][10]int

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			pixel[x][y] = 0
		}
	}
	fmt.Println(pixel)
}

// 向函数传递数组
func functransferArrToFunc() {
	arr := [20]int{1, 2, 3, 4, 5}
	arrSize := count(&arr)
	fmt.Println(arrSize)

	slice := arr[:]
	arrSize2 := count2(slice) //传递数组分片
	fmt.Println(arrSize2)
}

// 向函数传递数组，传递数组的指针
func count(arr *[20]int) int {
	index := 0
	for _, v := range *arr {
		if v == 0 {
			return index
		}
		index++
	}
	return index
}

// 向函数传递分片
func count2(arr []int) int {
	index := 0
	for _, v := range arr {
		if v == 0 {
			return index
		}
		index++
	}
	return index
}

// slice 切片定义
func sliceDefine() {
	//创建指向arr切片
	arr := []string{"a", "b", "c", "d", "e"}
	slice1 := arr[:]
	slice2 := arr[1:3]
	slice3 := arr[4:]
	fmt.Printf("%s, cap=%d, len=%d \n", slice1, cap(slice1), len(slice1)) // [a b c d e], cap=5, len=5
	fmt.Printf("%s, cap=%d, len=%d \n", slice2, cap(slice2), len(slice2)) // [b c], cap=4, len=2
	fmt.Printf("%s, cap=%d, len=%d \n", slice3, cap(slice3), len(slice3)) // [e], cap=1, len=1

	//直接创建分片
	slice4 := make([]int, 10, 100) //cap=10, len=100
	println(len(slice4), cap(slice4))
	for i := 0; i < len(slice4); i++ {
		slice4[i] = rand.Int()
	}

	fmt.Println("cap=%d, len=%d, %s \n", cap(slice4), len(slice4), slice4)

	// 直接创建切片，仅指定 cap
	s2 := make([]int, 100)        // 指定 cap
	fmt.Println(len(s2), cap(s2)) // 100, 100
}
