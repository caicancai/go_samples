package main

import "fmt"

//go基本类型

// 常量
const PI = 3.141592

// 常量枚举
const (
	Unknown = 0
	Female  = 1
	Male
)

// 常量枚举2
const beef, two, c = "eat", 2, "veg"

// iota使用
const (
	Sunday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

var (
	user = "hangc"
	flag = true
	city string
)

func main() {
	test4()
	test5()
	test6()
	//println("user = ", user)
	//test7()
}

// 类型转换
func test4() {
	const NUM int = 12
	var r1 = float32(NUM) + PI
	println(r1)
	fmt.Printf("%s,%t", user, flag)
	var num2 int8 = 12
	println(num2)
}

// 测试 iota枚举
func test5() {
	fmt.Printf("%d,%d,%d,%d,%d,%d,%d\n", Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}

// 使用函数内局部变量
func test6() {
	println("user = ", user)
	user := "vancy"
	println("test6.user = ", user)

	country := "guangzhou" // 局部变量初始化省略写法
	scope := uint16(10240) // 局部变量初始化省略写法，同时赋值类型
	println("country: ", country, ", scope: ", scope)
}

// 变量引用类型和值类型
func test7() {
	var n1 = 233
	var n2 = n1
	println(&n1, " ", &n2, " ", &n1 == &n2)
	// 值拷贝
	var s1 = "hello"
	var s2 = s1
	println(&s1, " ", &s2, " ", &s1 == &s2)
	var p1 *string = &s1
	fmt.Printf("&s1=%p, p1=%p \n", &s1, p1)
	println("&s1 == p1 : ", &s1 == p1)
	println("s1 == *p1 : ", s1 == *p1, " *p1=", *p1)
}
