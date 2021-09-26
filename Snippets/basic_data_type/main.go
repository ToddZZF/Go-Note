package main

import (
	"fmt"
	"math"
)

func main() {
	// 十进制
	var a int = 10
	fmt.Printf("打印十进制：%d \n", a) // 10
	fmt.Printf("打印二进制：%b \n", a) // 1010  占位符%b表示二进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("打印八进制：%o \n", b) // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("打印十六进制（小写）：%x \n", c) // ff
	fmt.Printf("打印十六进制（大写）：%X \n", c) // FF

	// 特殊技巧
	o := 0666
	fmt.Printf("特殊技巧打印：%d %[1]o %#[1]o\n", o) // "438 666 0666"
	x := int64(0xdeadbeef)
	fmt.Printf("特殊技巧打印：%d %[1]x %#[1]x %#[1]X\n", x)
	// Output:
	// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF

	// 浮点数
	fmt.Printf("打印浮点数：%f\n", math.Pi)
	fmt.Printf("打印指定位数的浮点数：%.2f\n", math.Pi)

	// 复数
	var c1 complex64 = 1 + 2i
	var c2 complex128 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)

	// 多行字符串
	s1 := `第\n一行
第\t二行
第\\三行
`
	fmt.Println(s1)

	// byte 或 rune
	traversalString()
}

func traversalString() {
	s := "Hello, 世界"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}
