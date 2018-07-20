package main

import (
	"fmt"
	"strconv"
	"time"
)

type Watch interface {
	name() string           //名稱
	size() string           //大小
	useTime() time.Duration //待機時間
	color() string          //顏色
}

type IWatch struct {
	version, watchColor string
	width, height       int
	time                time.Duration
}

func (i IWatch) name() string {
	return i.version
}

func (i IWatch) size() string {
	return strconv.Itoa(i.width) + " mm * " + strconv.Itoa(i.height) + " mm"
}

func (i IWatch) useTime() time.Duration {
	return i.time * time.Hour
}

func (i IWatch) color() string {
	return i.watchColor
}

func showProduce(w Watch) {
	fmt.Printf("手錶名稱：%s \n", w.name())
	fmt.Printf("大小：%s \n", w.size())
	fmt.Printf("待機時間：%v \n", w.useTime())
	fmt.Printf("顏色：%s \n\n", w.color())
}

func main() {
	IWatch1 := IWatch{version: "APPLE Watch 3", watchColor: "紫色", width: 38, height: 42, time: 80}
	IWatch2 := IWatch{version: "APPLE Watch Series 3", watchColor: "藍色", width: 40, height: 42, time: 120}
	showProduce(IWatch1)
	showProduce(IWatch2)
}
