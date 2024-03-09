package main

import (
	"REST/pkg"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/input", getInput)
	r.GET("/tpl", getTpl)
	r.GET("/", getHome)
	r.POST("/", postHome)
	//askAPI()
	askAPIBooking("las vegas")
	//askAPIPost("las vegas")
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

func askAPI() {
	url := "https://flight-fare-search.p.rapidapi.com/v2/flights/?from=LHR&to=DXB&date=%3CREQUIRED%3E&adult=1&type=economy&currency=USD"

	api := os.Getenv("API")
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-RapidAPI-Key", api)
	req.Header.Add("X-RapidAPI-Host", "flight-fare-search.p.rapidapi.com")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	v := gin.H{}
	err = json.NewDecoder(res.Body).Decode(&v)
	if err != nil {
		fmt.Errorf("%w\n", err)
		return
	}
	fmt.Printf("Body: %#v\n", v)
}
func askAPIPost(query string) {
	url := "https://tourist-attraction.p.rapidapi.com/typeahead"
	api := os.Getenv("API")
	payload := strings.NewReader(fmt.Sprintf("q=%s&language=en_US", query))
	req, err := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-RapidAPI-Key", api)
	req.Header.Add("X-RapidAPI-Host", "tourist-attraction.p.rapidapi.com")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	v := pkg.Typeahead{}
	err = json.NewDecoder(res.Body).Decode(&v)
	if err != nil {
		fmt.Errorf("%w\n", err)
		return
	}
	fmt.Printf("Body: %s\n", v.Results.Data[0].ResultObject.LocationString)
}

func askAPIBooking(query string) {
	url := fmt.Sprintf("https://booking-com15.p.rapidapi.com/api/v1/hotels/searchDestination?query=%s", query)

	api := os.Getenv("API")
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", api)
	req.Header.Add("X-RapidAPI-Host", "booking-com15.p.rapidapi.com")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	v := gin.H{}
	err = json.NewDecoder(res.Body).Decode(&v)
	if err != nil {
		fmt.Errorf("%w\n", err)
		return
	}
	fmt.Printf("Body booking: %v\n", v)
}
