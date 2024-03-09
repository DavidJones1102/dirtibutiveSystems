package main

import (
	"REST/pkg"
	_ "embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//go:embed templates/input.html
var input []byte

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/input", getInput)
	r.GET("/tpl", getTpl)
	r.GET("/", getHome)
	r.POST("/", postHome)
	r.Run("127.0.0.1:8080")
}

func getHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func postHome(c *gin.Context) {
	fmt.Printf("Query: %s\n", c.Query("id"))
	Body := pkg.FormA{}
	if errA := c.ShouldBind(&Body); errA != nil {
		fmt.Printf("Error body %s", Body.Foo)
	}
	fmt.Printf("Body %s\n", Body.Foo)
}
func getInput(c *gin.Context) {
	c.HTML(http.StatusOK, "input.html", nil)
}
func getTpl(c *gin.Context) {
	m := gin.H{"Message": "Hej z template"}
	c.HTML(http.StatusOK, "response.html", m)
}
