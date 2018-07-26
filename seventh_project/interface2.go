package main

import (
	"fmt"
)

type Phone interface {
	Sensor() string
	Monitor() string
	Button() string
}

type IPhone struct {
	varsion      string
	userName     string
	lockButton   string
	unlockButton string
}

type HTC struct {
	varsion      string
	userName     string
	lockButton1  string
	lockButton2  string
	unlockButton string
}

func (i IPhone) Sensor() string {
	return i.lockButton
}

func (i IPhone) Monitor() string {
	return i.userName + " 歡迎使用 IPhone " + i.varsion
}

func (i IPhone) Button() string {
	return i.unlockButton
}

func (h HTC) Sensor() string {
	return h.lockButton1 + "+" + h.lockButton2
}

func (h HTC) Monitor() string {
	return h.userName + " 歡迎使用 HTC " + h.varsion
}

func (h HTC) Button() string {
	return h.unlockButton
}

func Unlock(p Phone) {
	println("已使用 " + p.Sensor() + " 解鎖")
	println("Hi~ " + p.Monitor())
}

func Lock(p Phone) {
	println("已用 " + p.Button() + " 上鎖")
}

func main() {
	IPhone1 := IPhone{"X", "wally", "人臉辨識", "*"}
	IPhone2 := IPhone{"8", "Jenny", "指紋", "*"}
	iPhoneProcess(IPhone1)
	iPhoneProcess(IPhone2)

	HTC1 := HTC{"U11", "David", "虹膜辨識", "指紋", "*"}
	HTC2 := HTC{"M8", "Duke", "#", "7", "*"}
	HTCProcess(HTC1)
	HTCProcess(HTC2)
}

func iPhoneProcess(usePhone IPhone) {
	fmt.Print("\n")
	fmt.Printf("IPhone " + usePhone.varsion + " 鎖定解鎖流程：\n")
	Lock(usePhone)
	Unlock(usePhone)
}

func HTCProcess(usePhone HTC) {
	fmt.Print("\n")
	fmt.Printf("HTC " + usePhone.varsion + " 鎖定解鎖流程：\n")
	Lock(usePhone)
	Unlock(usePhone)
}
