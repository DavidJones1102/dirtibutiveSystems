package pkg

import "time"

type Typeahead struct {
	Results struct {
		Data []struct {
			ResultObject struct {
				Name        string `json:"name"`
				Latitude    string `json:"latitude"`
				Longitude   string `json:"longitude"`
				Description string `json:"description"`
				Photo       struct {
					Images struct {
						Original Photo `json:"original"`
					} `json:"images"`
				} `json:"photo"`
			} `json:"result_object"`
			Scope string `json:"scope"`
		} `json:"data"`
	} `json:"results"`
}

type Photo struct {
	Width  string `json:"width"`
	Height string `json:"height"`
	Url    string `json:"url"`
}

type Body struct {
	City      string
	StartDate time.Time
	EndDate   time.Time
}

type BookingResponse struct {
	Result []struct {
		ReviewScoreWord string   `json:"review_score_word"`
		ReviewScore     *float64 `json:"review_score"`
		HotelName       string   `json:"hotel_name"`
		MaxPhotoUrl     string   `json:"max_photo_url"`
	} `json:"result"`
}
