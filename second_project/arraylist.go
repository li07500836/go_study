package main

import (
	"fmt"
)

func main() {
	//宣告5筆測驗成績
	var x [5]int
	x[0] = 90
	x[1] = 80
	x[2] = 55
	x[3] = 48
	x[4] = 87
	//印出成績
	fmt.Printf("成績：")
	fmt.Println(x)

	//計算成績總和
	total := 0
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	//印出總和
	fmt.Printf("總和：%d\n", total)

	//取得平均數
	var average float32
	average = float32(total) / float32(len(x))
	fmt.Printf("平均數： %f\n", average)

	//作業2
	z := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	min := minResult(z)
	fmt.Printf("最小值： %d\n", min)
}

//計算最小值fun
func minResult(z []int) int {
	//預設最小值為陣列第一位
	min := int(z[0])
	//計算最小值
	for j := 0; j < len(z); j++ {
		if z[j] < min {
			min = z[j]
		}
	}
	return min
}
