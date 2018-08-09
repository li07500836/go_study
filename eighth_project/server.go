package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/parm/show/:watchid", getIWatchStrcut)
	e.GET("/phone/:watchid", getIWatch)
	e.Logger.Fatal(e.Start(":4356"))
}

func getIWatchStrcut(c echo.Context) error {
	watchid := c.Param("watchid")
	color := c.QueryParam("color")
	result := "watchid:" + watchid + "color:" + color
	return c.String(http.StatusOK, result)
}

// API 1
func getIWatch(c echo.Context) error {
	IWatch1 := IWatch{version: "APPLE Watch " + c.Param("watchid"), watchColor: c.QueryParam("color"), width: 38, height: 42}
	//IWatch2 := IWatch{version: "APPLE Watch Series 3", watchColor: "藍色", width: 40, height: 42}
	result := ShowProduce(IWatch1)
	//ShowProduce(IWatch2)
	return c.String(http.StatusOK, result)
}

type Watch interface {
	name() string  //名稱
	size() string  //大小
	color() string //顏色
}

type IWatch struct {
	version, watchColor string
	width, height       int
}

func (i IWatch) name() string {
	return i.version
}

func (i IWatch) size() string {
	return strconv.Itoa(i.width) + " mm * " + strconv.Itoa(i.height) + " mm"
}

func (i IWatch) color() string {
	return i.watchColor
}

func ShowProduce(w Watch) string {
	fmt.Printf("手錶名稱：%s \n", w.name())
	msgName := "手錶名稱:" + w.name()
	fmt.Printf("大小：%s \n", w.size())
	msgSize := "大小:" + w.size()
	fmt.Printf("顏色：%s \n\n", w.color())
	msgColor := "顏色:" + w.color()
	msg := msgName + ", " + msgSize + ", " + msgColor
	return msg
}
