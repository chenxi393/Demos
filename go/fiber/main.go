
package main
 
import (
        "fmt"
        "time"
 
        "github.com/gofiber/fiber"
)
 
//for (( c=1; c<=10; c++ )); do   curl "http://127.0.0.1:5555/$c";   done

func main() {
        app := fiber.New()
 
        app.Get("/:number", func(c *fiber.Ctx) {
                number := c.Params("number")
                go myfunc(number)
                c.Send(number)
        })
        app.Listen(5555)
}
func myfunc(number string) {
        fmt.Printf("number is %s \n", number)
        time.Sleep(1 * time.Second)
        fmt.Printf("number is now %s \n", number)
}