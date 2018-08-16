package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/weather", getWeather)
	e.Logger.Fatal(e.Start(":1111"))
}

func getWeather(a echo.Context) error {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return err
	}
	defer c.Close()

	//判斷是否存在redis裡
	isKeyExit, err := redis.Bool(c.Do("EXISTS", "weather"))
	if err != nil {
		return err
	}

	weather := ""
	if isKeyExit {
		//存在,取出資料回傳
		weather, err = redis.String(c.Do("GET", "weather"))
		if err != nil {
			fmt.Println("redis get failed:", err)
		}
	} else {
		//取DB資料
		weather = getDB()
		if weather == "" {
			//DB沒資料直接取API
			weather = getAPI()
			//寫入DB
			setDB(weather)
			//寫入Redis
			setRedis(weather)
		}
	}

	return a.String(http.StatusOK, weather)
}

// DB基本模型的定义
type Weather struct {
	ID      uint `gorm:"primary_key"`
	Weather string
}

//取DB資料
func getDB() string {
	//不存在連線DB取資料
	db, err := gorm.Open("mysql", "li07500836:qwe123@tcp(127.0.0.1:3306)/Weather?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接数据库失败")
	}
	defer db.Close()

	var weatherResult Weather
	if err := db.Table("Weather").First(&weatherResult); err != nil {
		fmt.Println("取資料庫資料失败")
		return ""
	}

	weather := weatherResult.Weather
	return weather

	return ""
}

//取API資料
func getAPI() string {
	resp, err := http.Get("http://weather.json.tw/api")
	if err != nil {
		// handle error
		fmt.Println("API連線失敗失败")
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("取API資料異常")
		return ""
	}

	fmt.Println(string(body))

	return string(body)
}

//寫入DB
func setDB(body string) {
	//不存在連線DB取資料
	db, err := gorm.Open("mysql", "li07500836:qwe123@tcp(127.0.0.1:3306)/Weather?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接数据库失败")
	}
	defer db.Close()

	var weather = Weather{ID: 1, Weather: body}
	db.Create(&weather)
}

//寫入Redis
func setRedis(body string) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
	}
	defer c.Close()

	_, err = c.Do("SET", "weather", body)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}
