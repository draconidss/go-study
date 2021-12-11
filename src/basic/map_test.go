package main

import (
	"fmt"
	"testing"
)

func TestMapDeclare(t *testing.T) {
	// 第一种方法
	var scores map[string]int = map[string]int{"english": 80, "chinese": 85}
	fmt.Println(scores)

	//第二种方法
	scores_2 := map[string]int{"english": 80, "chinese": 86}
	fmt.Println(scores_2)

	//第三种方法
	scores_3 := make(map[string]int)
	fmt.Println(scores_3)
}

func TestMapMethod(t *testing.T) {
	scores := make(map[string]int)
	//赋值，读取
	scores["chinese"] = 100
	fmt.Println(scores)
	fmt.Println(scores["chinese"]) //100
	//返回对应数据类型的默认值
	fmt.Println(scores["noExist"]) //0

	//遍历key和value
	fmt.Println("遍历key和value")
	for subject, score := range scores {
		fmt.Printf("key: %s, value: %d\n\n", subject, score)
	}

	//遍历key
	fmt.Println("遍历key")
	for subject := range scores {
		fmt.Printf("key: %s\n", subject)
	}

	//遍历value要用占位符
	fmt.Println("遍历value")
	for _, score := range scores {
		fmt.Printf("key: %d\n", score)
	}

	//删除
	delete(scores, "chinese")
	fmt.Println(scores["chinese"])

	//判断存在
	//其实字典的下标读取可以返回两个值，使用第二个返回值都表示对应的 key 是否存在，
	//若存在ok为true，若不存在，则ok为false
	if math, ok := scores["math"]; ok {
		fmt.Printf("math 的值是: %d", math)
	} else {
		fmt.Println("math 不存在")
	}

}

func main() {

}
