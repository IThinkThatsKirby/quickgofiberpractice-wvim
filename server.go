package main

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)
func main(){
	//instance of Fiber w/BLAZINGLY fast JSON marshals
	app:=fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	//routes
	app.Get("/",res)
	//Open a port for HTTP reqeusts
	app.Listen(":5000")
}
//struct for the request JSON
type Transaction struct {
	Payer string `json:"payer"`
	Points int `json:"points"`
	Timestamp string `json:"timestamp"`
}

//Handler JSON unmarsheling

func res(c *fiber.Ctx)error{
	data := new(Transaction)
	if err := c.BodyParser(data); err != nil{
		return err}
		println(data.Payer)
		println(data.Points)
		println(data.Timestamp)

	return c.SendString("Howdy world!")
}
