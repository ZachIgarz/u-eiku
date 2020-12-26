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
	sayHi  = "Hola!"
	sayBye = "Bye..!"
)

func hello(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			// Write

			//TODO: validar si es saludoo despedida

			//err := websocket.Message.Send(ws, "Hello, Client!")

			// Read
			msg := ""
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
			}

			sendMesaje(ws, c, msg)

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

		return "Buenas Noches UwU"
	}

	if number > 17 && number < 20 {
		return "Buenas tardes"
	}
	if number > 5 && number < 17 {

		return "Buenos dias"
	}

	return ""
}

func sendMesaje(ws *websocket.Conn, c echo.Context, receivedrMesage string) {

	var mesage string
	mesage = sayHi + " " + sendMesajeDayOrNight()
	//TODO: validar que el mensaje contenga palabras clave
	if receivedrMesage == "Bye!" {

		mesage = sayBye + " " + sendMesajeDayOrNight()
	}

	err := websocket.Message.Send(ws, mesage)

	if err != nil {
		c.Logger().Error(err)
	}
}
