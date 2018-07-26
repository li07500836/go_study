package main

import (
	"./Rocket"
)

type RocketReal struct {
	rocketName string
	startTime  int
}

func (r RocketReal) Name() string {
	return r.rocketName
}

func (r RocketReal) PushTime() int {
	return r.startTime
}

func main() {
	Rocket1 := RocketReal{"阿波羅號", 10}
	Rocket2 := RocketReal{"黑帝斯號", 5}

	Rocket.Launch(Rocket1)
	Rocket.Launch(Rocket2)
}
