package main

import (
	"fmt"
	"testing"
	"unsafe"
)

//整型
func TestInteger(t *testing.T) {
	var (
	/*		i1 int
			i2 int8
			i3 int16
			i4 int32
			i5 int64
			i6 uint
			i7 uint8
			i8 uint16
			i9 uint32
			i10 uint64*/
	)

	//二进制，以0b或0B为前缀
	var num01 int = 0b1100

	//八进制，以0o或0O为前缀
	var num02 int = 0o14

	//十六进制，以0x或0X为前缀
	var num03 int = 0xC

	fmt.Println("整数型示例")
	fmt.Printf("2进制数 %b 表示的是: %d \n", num01, num01)
	fmt.Printf("8进制数 %o 表示的是: %d \n", num02, num02)
	fmt.Printf("16进制数 %X 表示的是: %d \n", num03, num03)

}

//浮点型
func TestFloat(t *testing.T) {
	var myfloat01 float32 = 100000182
	var myfloat02 float32 = 100000187

	//精度有限
	fmt.Println()
	fmt.Println("浮点型示例")
	fmt.Println("float: ", myfloat01)
	fmt.Println("float: ", myfloat01+5)
	fmt.Println(myfloat02 == myfloat01+5)
}

//byte示例
/*byte，占用1个节字，就 8 个比特位（2^8 = 256，因此 byte 的表示范围 0->255），所以它和 uint8 类型本质上没有区别，它表示的是 ACSII 表中的一个字符。*/
func TestByte(t *testing.T) {
	var a byte = 65
	// 8进制写法: var a byte = '\101'     其中 \ 是固定前缀
	// 16进制写法: var a byte = '\x41'    其中 \x 是固定前缀

	var b uint8 = 66
	fmt.Println()
	fmt.Println("byte示例")
	fmt.Printf("a 的值: %c \nb 的值: %c\n", a, b)

	// 或者使用 string 函数
	// fmt.Println("a 的值: ", string(a)," \nb 的值: ", string(b))
}

//rune示例
/*rune，占用4个字节，共32位比特位，所以它和 uint32 本质上也没有区别。它表示的是一个 Unicode字符（Unicode是一个可以表示世界范围内的绝大部分字符的编码规范）。*/
func TestRune(t *testing.T) {
	var a byte = 'A'
	var b rune = '中'
	fmt.Println()
	fmt.Println("rune示例")
	fmt.Printf("a 占用 %d 个字节数\nb 占用 %d 个字节数\n", unsafe.Sizeof(a), unsafe.Sizeof(b))
}

func TestString(t *testing.T) {
	var mystr01 string = "hello,中国"
	//本质上string是一个字符数组
	var mystr02 [5]byte = [5]byte{104, 101, 108, 108, 111}
	var myStr03 string = `\n\r`
	fmt.Println()
	fmt.Println("string示例")
	fmt.Printf("mystr01: %s\n", mystr01)
	fmt.Printf("mystr02: %s\n", mystr02)
	fmt.Printf("mystr03: %s\n", myStr03)
	fmt.Printf("还原的mystr03: %q\n", myStr03)
	//utf-8中，英文占1个字节。中文占3个字节
	fmt.Println("占用字节: ", len(mystr01))
}
