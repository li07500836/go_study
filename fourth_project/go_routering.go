package main

import (
	"fmt"
	"sync"
	"time"
)

type factory struct {
	name    string  //姓名
	produce int     //一次產生的數量
	need    float32 //一次需要的材料
	wait    int     //每次吃完後要休息的時間
	color   int     //訊息顯示的顏色
}

func main() {
	//紙漿從0開始
	var pulp float32
	totalPulp := 0
	//產生的紙
	paperCount := 0
	totalPaper := 0
	//印刷的數量
	printPaper := 60000

	//伐木工場
	wood1 := factory{
		name:    "伐木工場1",
		produce: 1,
		need:    0,
		wait:    1,
		color:   31,
	}

	//造紙工廠
	paper1 := factory{
		name:    "造紙工場1",
		produce: 5000,
		need:    0.5,
		wait:    1,
		color:   32,
	}
	paper2 := factory{
		name:    "造紙工場2",
		produce: 3000,
		need:    0.3,
		wait:    1,
		color:   33,
	}

	//印刷工廠
	print1 := factory{
		name:    "印刷工廠1",
		produce: 6000,
		need:    6000,
		wait:    1,
		color:   35,
	}

	//建立WaitGroup
	var wg sync.WaitGroup
	//建立互斥鎖
	var mutex1 sync.Mutex
	var mutex2 sync.Mutex
	var mutex3 sync.Mutex
	//總共100跟樹木

	wg.Add(1)

	//使用goroutine讓工場開始工作
	//伐木工廠運作
	go wood(wood1, &pulp, &totalPulp, &wg, mutex1)
	//造紙工廠運作
	go paper(paper1, &pulp, &paperCount, &totalPaper, &wg, mutex2)
	go paper(paper2, &pulp, &paperCount, &totalPaper, &wg, mutex2)
	//印刷工廠運作
	go print(print1, &paperCount, &printPaper, &wg, mutex3)
	//等待所有工作完成
	wg.Wait()

	//所有工作都完成後才會繼續執行程式碼
	fmt.Printf("%c[0;40;36m已完成所有工作%c[0m\n", 0x1B, 0x1B)
	fmt.Printf("%s 總共產生 %d 頓紙漿\n", wood1.name, totalPulp)
	fmt.Printf("%s 和 %s 總共產生 %d 張紙\n", paper1.name, paper2.name, totalPaper)
}

//製造紙漿
func wood(x factory, pulp *float32, totalPulp *int, wg *sync.WaitGroup, mutex sync.Mutex) {

	//一直產生紙漿
	for i := *pulp; i >= 0; i = *pulp {
		//數量上鎖
		mutex.Lock()
		*pulp = *pulp + float32(x.produce)
		*totalPulp = *totalPulp + x.produce
		fmt.Printf("%c[0;40;%dm%s 產生 %d 頓紙漿，目前有 %f 頓紙漿！%c[0m\n", 0x1B, x.color, x.name, x.produce, *pulp, 0x1B)

		//數量解鎖
		mutex.Unlock()

		time_s := time.Duration(x.wait) * time.Second
		time.Sleep(time_s)
	}
}

//製造紙漿
func paper(x factory, pulp *float32, paperCount *int, totalPaper *int, wg *sync.WaitGroup, mutex sync.Mutex) {

	//一直產生紙
	for i := *paperCount; i >= 0; i = *paperCount {
		if *pulp > x.need {
			//數量上鎖
			mutex.Lock()
			*pulp = *pulp - x.need
			*paperCount = *paperCount + x.produce
			*totalPaper = *totalPaper + x.produce
			fmt.Printf("%c[0;40;%dm%s 產生 %d 張紙，目前有 %d 張紙！%c[0m\n", 0x1B, x.color, x.name, x.produce, *paperCount, 0x1B)

			//數量解鎖
			mutex.Unlock()
		} else {
			fmt.Printf("%c[0;40;%dm%s 無紙漿可使用！%c[0m\n", 0x1B, x.color, x.name, 0x1B)
		}
		time_s := time.Duration(x.wait) * time.Second
		time.Sleep(time_s)
	}
}

//印刷
func print(x factory, paperCount *int, printPaper *int, wg *sync.WaitGroup, mutex sync.Mutex) {
	var print int

	//在印刷完成前不會停止工作
	for i := *printPaper; i > 0; i = *printPaper {
		if float32(*paperCount) > x.need {
			//樹木數量上鎖
			mutex.Lock()

			if *printPaper >= x.produce {
				*paperCount = *paperCount - x.produce
				*printPaper = *printPaper - x.produce
				print = x.produce
			} else {
				print = *printPaper
				*paperCount = *paperCount - print
				*printPaper = 0
			}
			fmt.Printf("%c[0;40;%dm%s 印刷 %d 張紙，剩下 %d 張需要印刷！%c[0m\n", 0x1B, x.color, x.name, print, *printPaper, 0x1B)

			//樹木數量解鎖
			mutex.Unlock()
		} else {
			fmt.Printf("%c[0;40;%dm%s 無紙可印刷！%c[0m\n", 0x1B, x.color, x.name, 0x1B)
		}
		time_s := time.Duration(x.wait) * time.Second
		time.Sleep(time_s)
	}

	//工作完成 回報WaitGroup -1
	wg.Done()
}
