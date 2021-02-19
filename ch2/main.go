package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 变量声明
	var i int = 1
	fmt.Println("i = ", i)

	// 变量类型推导
	var j = 10
	fmt.Println("j = ", j)

	// 一次声明多个变量
	var (
		m int = 100
		n     = 1000
	)
	fmt.Println("m =", m, ", n =", n)

	// 简短声明
	si := 10
	sb := false
	ss := "Hello"
	fmt.Println("si =", si, ", sb =", sb, ", ss = ", ss)

	// 浮点数声明
	var f32 float32 = 1.23
	var f64 float64 = 12.3456
	fmt.Println("f32 =", f32, ", f64 =", f64)

	// 数字类型强制转换
	i2f := float64(i)
	fmt.Printf("i = %v, i2f type = %T, value = %v\n", i, i2f, i2f)
	// i = 1, i2f type = float64, value = 1
	f2i := int(f64)
	fmt.Printf("f64 = %v, f2i type = %T, value = %v\n", f64, f2i, f2i)
	// f64 = 12.3456, f2i type = int, value = 12

	// 布尔型声明
	var bf bool = false
	var bt bool = true
	fmt.Println("bf =", bf, ", bt =", bt)

	// 字符串声明
	var s1 string = "Hello"
	var s2 string = "世界"
	fmt.Println("s1 =", s1, ", s2 =", s2)

	// 操作符与字符串
	fmt.Println("s1 + s2 = ", s1+s2)

	// 字符串和整数互转
	i2s := strconv.Itoa(i)
	fmt.Printf("i2s type = %T, value = %v\n", i2s, i2s)
	// i2s type = string, value = 1
	if s2i, err := strconv.Atoi(i2s); err == nil {
		fmt.Printf("s2i type = %T, value = %v\n", s2i, s2i)
	}
	// s2i type = int, value = 1

	// 字符串和浮点数互转
	vs := "3.1415926535"
	if s2f32, err := strconv.ParseFloat(vs, 32); err == nil {
		fmt.Printf("s2f32 type = %T, value = %v\n", s2f32, s2f32)
	}
	// s2f32 type = float64, value = 3.1415927410125732
	if s2f64, err := strconv.ParseFloat(vs, 64); err == nil {
		fmt.Printf("s2f64 type = %T, value = %v\n", s2f64, s2f64)
	}
	// s2f64 type = float64, value = 3.1415926535

	vf := 3.1415926535
	s32 := strconv.FormatFloat(vf, 'f', -1, 32)
	fmt.Printf("s32 type = %T, value = %v\n", s32, s32)
	// s32 type = string, value = 3.1415927
	s64 := strconv.FormatFloat(vf, 'f', -1, 64)
	fmt.Printf("s64 type = %T, value = %v\n", s64, s64)
	// s64 type = string, value = 3.1415926535

	// 字符串和布尔类型互转
	vbs := "true"
	if s2b, err := strconv.ParseBool(vbs); err == nil {
		fmt.Printf("s2b type = %T, value = %v\n", s2b, s2b)
	}
	// s2b type = bool, value = true
	vb := true
	b2s := strconv.FormatBool(vb)
	fmt.Printf("b2s type = %T, value = %v\n", b2s, b2s)
	// b2s type = string, value = true

	// 零值
	var (
		zi int
		zf float64
		zb bool
		zs string
	)
	fmt.Println("zi =", zi, ", zf =", zf, ", zb =", zb, ", zs =", zs)

	// 指针
	pi := &i
	fmt.Println("pi =", pi, ", *pi =", *pi, ", i =", i)

	// 常量定义
	const cs1 string = "Hello"
	const cs2 = "世界"
	const (
		cs3 string = "你好"
		cs4        = "World"
	)
	fmt.Println("cs1 =", cs1, ", cs2 =", cs2, ", cs3 =", cs3, ", cs4 =", cs4)

	// iota
	const (
		one   = 1
		two   = 2
		three = 3
		four  = 4
	)
	fmt.Println("one =", one, ", two =", two, ", three =", three, ", four =", four)
	// one = 1 , two = 2 , three = 3 , four = 4

	const (
		five = iota + 5
		six
		seven
		eight
	)
	fmt.Println("five =", five, ", six =", six, ", seven =", seven, ", eight =", eight)
	// five = 5 , six = 6 , seven = 7 , eight = 8
}
