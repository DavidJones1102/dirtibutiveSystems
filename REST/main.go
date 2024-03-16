package main

import (
	"REST/pkg"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"time"
)

var InternalAPI = "1234"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", getInput)
	r.POST("/", buildResponse)
	r.Run("127.0.0.1:8080")
}

func getInput(c *gin.Context) {
	c.HTML(http.StatusOK, "input.html", nil)
}

func buildResponse(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", "Cannot parse your form")
		return
	}
	inApi := c.PostForm("api")
	if inApi != InternalAPI {
		c.HTML(http.StatusBadRequest, "error.html", "Invalid api key")
		return
	}
	city := c.PostForm("city")
	startDate, validStart := time.Parse(time.DateOnly, c.PostForm("start_date"))
	endDate, validEnd := time.Parse(time.DateOnly, c.PostForm("end_date"))
	if validStart != nil || validEnd != nil || city == "" {
		c.HTML(http.StatusBadRequest, "error.html", "Invalid body")
		return
	}

	if startDate.After(endDate) {
		c.HTML(http.StatusBadRequest, "error.html", "End date should be after start date")
		return
	}

	if startDate.Before(time.Now()) {
		c.HTML(http.StatusBadRequest, "error.html", "You cannot travel in time ...")
		return
	}
	Body := pkg.Body{
		City:      city,
		StartDate: startDate,
		EndDate:   endDate,
	}
	typeahead := askAPIAttractions(Body.City)
	if len(typeahead.Results.Data) <= 0 {
		c.HTML(http.StatusOK, "error.html", "this place doesn't exist (at least in our database)")
		return
	}
	bookingResponse := askAPIBookingHotel(
		typeahead.Results.Data[0].ResultObject.Latitude,
		typeahead.Results.Data[0].ResultObject.Longitude,
		Body.StartDate.Format(time.DateOnly),
		Body.EndDate.Format(time.DateOnly),
	)
	c.HTML(http.StatusOK, "response.html", gin.H{
		"Attr":   typeahead.Results.Data[0],
		"Hotels": bookingResponse.Result,
	})
}

func askAPIAttractions(query string) pkg.Typeahead {
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
	Typeahead := pkg.Typeahead{}
	err = json.NewDecoder(res.Body).Decode(&Typeahead)
	if err != nil {
		fmt.Errorf("%w\n", err)
		return pkg.Typeahead{}
	}
	return Typeahead
}

func askAPIBookingHotel(latitude, longitude, arrivalDate, departureDate string) pkg.BookingResponse {
	url := fmt.Sprintf(
		"https://booking-com.p.rapidapi.com/v1/hotels/search-by-coordinates?locale=en-gb&room_number=1&checkin_date=%s&checkout_date=%s&filter_by_currency=EUR&longitude=%s&latitude=%s&adults_number=2&order_by=popularity&units=metric&include_adjacency=true&children_ages=0",
		arrivalDate, departureDate, longitude, latitude,
	)

	api := os.Getenv("API")
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", api)
	req.Header.Add("X-RapidAPI-Host", "booking-com.p.rapidapi.com")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	hotelsList := pkg.BookingResponse{}
	_ = json.NewDecoder(res.Body).Decode(&hotelsList)
	return hotelsList
}
