package controller

import (
	"html/template"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	temp, err := template.ParseFiles("views/index.html")
	if err != nil {
		panic(err)
	}

	return temp.Execute(c.Response().BodyWriter(), nil)
}

func Login(c *fiber.Ctx) error{
	temp, err := template.ParseFiles("views/login.html")
	if err != nil{
		panic(err)
	}

	return temp.Execute(c.Response().BodyWriter(), nil)
}

func Register(c *fiber.Ctx) error{
	temp, err := template.ParseFiles("views/register.html")
	if err != nil{
		panic(err)
	}

	return temp.Execute(c.Response().BodyWriter(), nil)
}

func Account(c *fiber.Ctx) error{
	temp, err := template.ParseFiles("views/account.html")
	if err != nil{
		panic(err)
	}

	return temp.Execute(c.Response().BodyWriter(), nil)
}