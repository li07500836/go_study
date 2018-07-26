package Rocket

import (
	"fmt"
	"strconv"
)

type Rocket interface {
	Name() string  //火箭名稱
	PushTime() int //倒數時程
}

func Launch(rocket Rocket) {
	fmt.Println(rocket.Name() + " 開始倒數發射,倒數 " + strconv.Itoa(rocket.PushTime()) + " 秒")
	for i := rocket.PushTime(); i >= 0; i-- {
		fmt.Printf("%d\n", i)
	}
	fmt.Printf("%s 已升空 \n\n", rocket.Name())
}
