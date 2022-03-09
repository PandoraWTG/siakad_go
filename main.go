package main

import(
  "github.com/gofiber/fiber/v2"
)

func main(){
  route := fiber.New()

  route.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Testing")
  })
  route.Listen(":3000")
}
