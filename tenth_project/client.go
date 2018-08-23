package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	//一秒執行1次
	c.AddFunc("*/1 * * * * *", func() { goroute() })
	//goroute()
	c.Start()
}

func goroute() {
	count := 10

	//建立WaitGroup
	var wg sync.WaitGroup
	//建立互斥鎖
	var mutex sync.Mutex

	wg.Add(1)

	//使用goroutine讓工場開始工作
	//取API運作
	go getAPI(&count, &wg, mutex)
	//等待所有工作完成
	wg.Wait()

	//所有工作都完成後才會繼續執行程式碼
	fmt.Printf("%c[0;40;36m已完成所有工作%c[0m\n", 0x1B, 0x1B)
}

//連線API
func getAPI(count *int, wg *sync.WaitGroup, mutex sync.Mutex) {

	wait := 1
	//初始時間
	now := time.Now()
	//單筆最大/最小時間
	var maxAPI float32 = 0.00
	var minAPI float32 = 10.00
	allCount := *count
	//一秒call1次
	for i := *count; i >= 0; i = *count {
		//初始時間
		nowSub := time.Now()
		//遞減
		*count = *count - wait
		//數量上鎖
		mutex.Lock()

		resp, err := http.Get("http://weather.json.tw/api")
		if err != nil {
			// handle error
			fmt.Println("API連線失敗失败")
		}

		defer resp.Body.Close()
		//body, err := ioutil.ReadAll(resp.Body)
		//if err != nil {
		//	fmt.Println("取API資料異常")
		//}
		//印API結果
		//fmt.Println(string(body))
		//數量解鎖
		mutex.Unlock()

		nowSub2 := time.Now()
		subTime := nowSub2.Sub(nowSub)
		fmt.Println("單一連線花費時間：", subTime.Seconds(), "秒")
		if float32(subTime.Seconds()) > maxAPI {
			maxAPI = float32(subTime.Seconds())
		}
		if float32(subTime.Seconds()) < minAPI {
			minAPI = float32(subTime.Seconds())
		}

		//休息時間
		//time_s := time.Duration(wait) * time.Second
		//time.Sleep(time_s)
	}
	now2 := time.Now()
	subAllTime := now2.Sub(now)
	avTime := float32(subAllTime.Seconds()) / float32(allCount)
	fmt.Println("單一最大花費時間：", maxAPI, "秒")
	fmt.Println("單一最小花費時間：", minAPI, "秒")
	fmt.Println("壓測總數花費時間：", subAllTime.Seconds(), "秒")
	fmt.Println("平均花費時間", avTime, "秒")
	//工作完成 回報WaitGroup -1
	wg.Done()
}
