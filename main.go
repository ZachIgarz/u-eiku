package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
)

const (
	hi = "Hola"
)

func hello(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			// Write

			//err := websocket.Message.Send(ws, "Hello, Client!")
			err := websocket.Message.Send(ws, sendMesajeDayOrNight())
			if err != nil {
				c.Logger().Error(err)
			}

			// Read
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
			}
			fmt.Printf("%s\n", msg)
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "./public")
	e.GET("/ws", hello)
	e.Logger.Fatal(e.Start(":1323"))
}

//Escucha la hora del sistema
func listenSistemHour() string {

	return fmt.Sprint(time.Now().Hour())
}

func sendMesajeDayOrNight() string {
	actualTime := listenSistemHour()

	number, err := strconv.ParseUint(actualTime, 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	if number > 20 {

		return hi + " Buenas Noches UwU"
	}

	if number > 17 && number < 20 {
		return hi + " Buenas tardes"
	}

	return ""
}
