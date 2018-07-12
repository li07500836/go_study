package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)
	// goroutine1
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		//1.因程式沒等待go route都結束的機制,造成還在塞channel時,channel通道就已經被關閉
		//跑完寫入迴圈後,關閉channel通道
		close(ch)
	}()

	// goroutine2
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("作業1 執行ok", "\n")

	factorial()
}

//作業2
func factorial() {
	ch := make(chan int, 50)
	chCount := make(chan int, 50)
	now := time.Now()
	// goroutine1(計算階層)
	for num := 1; num < 21; num++ {
		go func() {
			sum := 1
			for i := 1; i <= num; i++ {
				sum = sum * i
			}
			ch <- sum
			chCount <- num
			fmt.Println(num, "! = ", sum, " 執行完畢！")
		}()
		time.Sleep(time.Second * 1)
	}

	// goroutine2（加總）
	go func() {
		totalSum := 0
		totalCount := 0
		for {
			a, ok := <-ch
			if ok {
				//加總
				totalSum = totalSum + a
			}
			//判斷線程數量
			b, check := <-chCount
			if check {
				totalCount = totalCount + 1
				b = b + 1
				//fmt.Println(b, "個線程執行完畢")
				if totalCount == 20 {
					fmt.Println("總加總為: ", totalSum)
					fmt.Println("全部線程執行完畢")
					now2 := time.Now()
					subTime := now2.Sub(now)
					fmt.Println("執行時間：", subTime.Seconds(), "秒")
					close(ch)
					close(chCount)
				}
			}
		}
	}()

	time.Sleep(time.Second * 30)
	fmt.Println("作業2 執行ok", "\n")
}
