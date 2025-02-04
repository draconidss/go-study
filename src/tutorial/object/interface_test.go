package object

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"testing"
	"text/tabwriter"
	"time"
)

// 接口定义
// 接口内嵌，和结构体内嵌相似
type InsidePhone interface {
	call_()
}

type Phone interface {
	call()
	// 接口内嵌
	InsidePhone
}

func TestUsage(t *testing.T) {
	//空接口的type和value都为nil
	var i interface{}
	fmt.Printf("type: %T, value: %v", i, i)

	//可以承载任意类型的值，也可以用空接口作为参数接收任意值
	// 存 int 没有问题
	i = 1
	fmt.Println(i)
	// 存字符串也没有问题
	i = "hello"
	fmt.Println(i)
	// 存布尔值也没有问题
	i = false
	fmt.Println(i)

	//定义一个接收任意类型的切片
	any := make([]interface{}, 5)
	any[0] = 11
	any[1] = "hello world"
	any[2] = []int{11, 22, 33, 44}
	for _, value := range any {
		fmt.Println(value)
	}
}

// 注意点
func TestInterfaceTip(t *testing.T) {
	//1.接口无法向下转型即使转成空接口和转出空接口的类型都是一样
	// 声明a变量, 类型int, 初始值为1
	var a int = 1
	// 声明i变量, 类型为interface{}, 初始值为a, 此时i的值变为1
	var i interface{} = a
	fmt.Println(i)
	// 声明b变量, 尝试赋值i
	//var b int = i
	//fmt.Println(b)

	//2.当空接口承载数组和切片后，该对象无法再进行切片
	//sli := []int{2, 3, 5, 7, 11, 13}
	//var i_1 interface{} = sli
	//g := i_1[1:3]
	//fmt.Println(g)

	//3.当你使用空接口来接收任意类型的参数时，它的静态类型是 interface{}，
	//但动态类型（是 int，string 还是其他类型）我们并不知道，因此需要使用类型断言。

}

type iPhone struct {
	name string
}

func (phone iPhone) call() {
	fmt.Println("Hello, iPhone.")
}

func (phone iPhone) call_() {
	fmt.Println("Hello, iPhone.")
}

func (phone iPhone) sendWechat() {
	fmt.Println("Hello, Wechat.")
}

// 测试实现接口的结构体方法的调用限制
func TestMethodExtendError(t *testing.T) {
	//这里声明类型为定义的Phone接口
	var phone Phone
	phone = iPhone{name: "ming's iphone"}
	phone.call()
	//无法调用该方法
	//phone.sendWechat()
}

// 修改为不显示的声明为Phone接口类型
func TestMethodExtendSuccess(t *testing.T) {
	phone := iPhone{name: "ming's iphone"}
	phone.call()
	phone.sendWechat()
}

type Nokia struct {
	name string
}

// 结构体实现接口方法，隐式，自动实现
// 一个类型如果定义了接口的所有方法，那它就隐式地实现了该接口。
func (phone *Nokia) call() {
	fmt.Println("我是Nokia打电话")
}

//--------------------------------------多态------------------------------------------

// 只要实现了这两个方法"就是一个商品",必须都实现
type Good interface {
	settleAccount() int
	orderInfo() string
}

type PhoneStruct struct {
	name     string
	quantity int
	price    int
}

type FreeGift struct {
	name     string
	quantity int
	price    int
}

// Phone
func (phone PhoneStruct) settleAccount() int {
	return phone.quantity * phone.price
}
func (phone PhoneStruct) orderInfo() string {
	return "您要购买" + strconv.Itoa(phone.quantity) + "个" +
		phone.name + "计：" + strconv.Itoa(phone.settleAccount()) + "元"
}

// FreeGift
func (gift FreeGift) settleAccount() int {
	return 0
}
func (gift FreeGift) orderInfo() string {
	return "您要购买" + strconv.Itoa(gift.quantity) + "个" +
		gift.name + "计：" + strconv.Itoa(gift.settleAccount()) + "元"
}

func calculateAllPrice(goods []Good) int {
	var allPrice int
	for _, good := range goods {
		fmt.Println(good.orderInfo())
		allPrice += good.settleAccount()
	}
	return allPrice
}

func TestMultiType(t *testing.T) {
	iPhone := PhoneStruct{
		name:     "iPhone",
		quantity: 1,
		price:    8000,
	}
	earphones := FreeGift{
		name:     "耳机",
		quantity: 1,
		price:    200,
	}

	//这里体现了多态
	goods := []Good{iPhone, earphones}
	allPrice := calculateAllPrice(goods)
	fmt.Printf("该订单总共需要支付 %d 元", allPrice)
}

// ==============================================扩展：sort 排序==============================================
// 参考：https://gopl-zh.github.io/ch7/ch7-06.html
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

// 自定义排序
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func TestSortByArtist(t *testing.T) {
	// 正序
	sort.Sort(byArtist(tracks))
	// 逆序
	sort.Sort(sort.Reverse(byArtist(tracks)))
	fmt.Printf("%#v\n", tracks)
}

// 使用内嵌函数来自定义排序
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func TestSortByFunc(t *testing.T) {
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})

}

// ============================================断言============================================

//Type Assertion（中文名叫：类型断言）
//1.仅能对静态类型为空接口（interface{}）的对象进行断言
//2.类型断言完成后，实际上会返回静态类型为你断言的类型的对象，而要清楚原来的静态类型为空接口类型（interface{}）

//通过它可以做到以下几件事情
//检查 i 是否为 nil
//检查 i 存储的值是否为某个类型
//通过类型断言查询接口

// 第一种：t := i.(T)
func TestAssertion_1(t *testing.T) {
	//var i interface{} = nil
	var i interface{} = 10
	t1 := i.(int)
	fmt.Println(t1)

	fmt.Println("=====分隔线=====")

	t2 := i.(string)
	//触发panic
	fmt.Println(t2)
}

// 第二种：t, ok:= i.(T)
// 如果断言失败不会触发panic，而是将ok设为false，此时t为T的零值
func TestTestAssertion_2(t *testing.T) {
	var i interface{} = 10
	t1, ok := i.(int)
	fmt.Printf("%d-%t\n", t1, ok)

	t2, ok := i.(float64)
	fmt.Printf("%f-%t\n", t2, ok)

	fmt.Println("=====分隔线2=====")

	var k interface{} // nil
	t3, ok := k.(interface{})
	fmt.Println(t3, "-", ok)

	fmt.Println("=====分隔线3=====")
	k = 10
	t4, ok := k.(interface{})
	fmt.Printf("%d-%t\n", t4, ok)

	t5, ok := k.(int)
	fmt.Printf("%d-%t\n", t5, ok)
}

// 区分多种类型,这里会隐式转型为接口类型
func findType(i interface{}) {
	//只能在switch中使用
	switch x := i.(type) {
	case int:
		fmt.Println(x, "is int")
	case string:
		fmt.Println(x, "is string")
	case nil:
		fmt.Println(x, "is nil")
	default:
		fmt.Println(x, "not type matched")
	}
}

// 主动转换为接口类型
func findTypeByInterface(i interface{}) {
	//主动转换
	findType(interface{}(i))
}

func TestTypeSwitch(t *testing.T) {
	findTypeByInterface(10)      // int
	findTypeByInterface("hello") // string

	var k interface{} // nil
	findTypeByInterface(k)

	findTypeByInterface(10.23) //float64
}

// 通过类型断言查询接口
// 我们想使用 WriteString 方法避免分配临时拷贝，但实现 io.Writer 的类型不一定有该方法
// 所以我们可以定义一个只有这个方法的新接口，并且使用类型断言来检测是否w的动态类型满足这个新接口，满足说明有 WriteString 那就调用，否则就用 Write 方法
// tips: 这里只是例子，标准库中有接口 io.WriteString
func writeString(w io.Writer, s string) (n int, err error) {
	type stringWriter interface {
		WriteString(string) (n int, err error)
	}
	// 如果断言类型成功
	if sw, ok := w.(stringWriter); ok {
		return sw.WriteString(s) // avoid a copy
	}
	return w.Write([]byte(s)) // allocate temporary copy
}
