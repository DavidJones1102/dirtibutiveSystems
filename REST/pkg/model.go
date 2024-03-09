package pkg

type FormA struct {
	Foo string `json:"foo"`
}

type Typeahead struct {
	Status  int    `json:"status"`
	Msg     string `json:"msg"`
	Results struct {
		Data []struct {
			ResultType   string `json:"result_type"`
			ResultObject struct {
				LocationId     string `json:"location_id"`
				Name           string `json:"name"`
				Latitude       string `json:"latitude"`
				Longitude      string `json:"longitude"`
				Timezone       string `json:"timezone"`
				LocationString string `json:"location_string"`
			} `json:"result_object"`
			Scope string `json:"scope"`
		} `json:"data"`
		PartialContent bool `json:"partial_content"`
	} `json:"results"`
}
