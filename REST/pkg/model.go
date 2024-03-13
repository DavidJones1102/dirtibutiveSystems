package pkg

import "time"

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
				Description    string `json:"description"`
				Photo          struct {
					Images struct {
						Small    Photo `json:"small"`
						Original Photo `json:"original"`
						Large    Photo `json:"large"`
					} `json:"images"`
				} `json:"photo"`
			} `json:"result_object"`
			Scope string `json:"scope"`
		} `json:"data"`
		PartialContent bool `json:"partial_content"`
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
		IsCityCenter        int         `json:"is_city_center"`
		MainPhotoUrl        string      `json:"main_photo_url"`
		Preferred           int         `json:"preferred"`
		CcRequired          int         `json:"cc_required"`
		UrgencyMessage      string      `json:"urgency_message,omitempty"`
		DistrictId          int         `json:"district_id"`
		PreferredPlus       int         `json:"preferred_plus"`
		HotelNameTrans      string      `json:"hotel_name_trans"`
		NativeAdsCpc        int         `json:"native_ads_cpc"`
		MinTotalPrice       float64     `json:"min_total_price"`
		ChildrenNotAllowed  interface{} `json:"children_not_allowed"`
		InBestDistrict      int         `json:"in_best_district"`
		Address             string      `json:"address"`
		DefaultWishlistName string      `json:"default_wishlist_name"`
		BlockIds            []string    `json:"block_ids"`
		Id                  string      `json:"id"`
		RibbonText          string      `json:"ribbon_text,omitempty"`
		Longitude           float64     `json:"longitude"`
		UpdatedCheckin      interface{} `json:"updated_checkin"`
		IsBeachFront        int         `json:"is_beach_front"`
		CountryTrans        string      `json:"country_trans"`
		CurrencyCode        string      `json:"currency_code"`
		DefaultLanguage     string      `json:"default_language"`
		Checkout            struct {
			From  string `json:"from"`
			Until string `json:"until"`
		} `json:"checkout"`
		District                 string  `json:"district"`
		MobileDiscountPercentage float64 `json:"mobile_discount_percentage"`
		IsMobileDeal             int     `json:"is_mobile_deal"`
		PriceIsFinal             int     `json:"price_is_final"`
		CityInTrans              string  `json:"city_in_trans"`
		ReviewNr                 *int    `json:"review_nr"`
		Bwallet                  struct {
			HotelEligibility int `json:"hotel_eligibility"`
		} `json:"bwallet"`
		Extended   int    `json:"extended"`
		CityNameEn string `json:"city_name_en"`
		Districts  string `json:"districts"`
		Badges     []struct {
			BadgeVariant string `json:"badge_variant"`
			Id           string `json:"id"`
			Text         string `json:"text"`
		} `json:"badges"`
		DistanceToCc            string      `json:"distance_to_cc"`
		ClassIsEstimated        int         `json:"class_is_estimated"`
		AccommodationTypeName   string      `json:"accommodation_type_name"`
		Cc1                     string      `json:"cc1"`
		UpdatedCheckout         interface{} `json:"updated_checkout"`
		Ufi                     int         `json:"ufi"`
		City                    string      `json:"city"`
		Soldout                 int         `json:"soldout"`
		AccommodationType       int         `json:"accommodation_type"`
		IsSmartDeal             int         `json:"is_smart_deal"`
		CompositePriceBreakdown struct {
			GrossAmount struct {
				AmountUnrounded string  `json:"amount_unrounded"`
				AmountRounded   string  `json:"amount_rounded"`
				Value           float64 `json:"value"`
				Currency        string  `json:"currency"`
			} `json:"gross_amount"`
			Benefits []struct {
				Kind         string      `json:"kind"`
				Identifier   string      `json:"identifier"`
				Icon         interface{} `json:"icon"`
				Details      string      `json:"details"`
				Name         string      `json:"name"`
				BadgeVariant string      `json:"badge_variant"`
			} `json:"benefits"`
			GrossAmountPerNight struct {
				Currency        string  `json:"currency"`
				AmountUnrounded string  `json:"amount_unrounded"`
				AmountRounded   string  `json:"amount_rounded"`
				Value           float64 `json:"value"`
			} `json:"gross_amount_per_night"`
			Items []struct {
				Details *string `json:"details"`
				Base    struct {
					BaseAmount int    `json:"base_amount,omitempty"`
					Kind       string `json:"kind"`
					Percentage int    `json:"percentage,omitempty"`
				} `json:"base"`
				InclusionType string `json:"inclusion_type,omitempty"`
				ItemAmount    struct {
					Currency        string  `json:"currency"`
					Value           float64 `json:"value"`
					AmountUnrounded string  `json:"amount_unrounded"`
					AmountRounded   string  `json:"amount_rounded"`
				} `json:"item_amount"`
				Name       string `json:"name"`
				Kind       string `json:"kind"`
				Identifier string `json:"identifier,omitempty"`
			} `json:"items"`
			AllInclusiveAmount struct {
				Value           float64 `json:"value"`
				AmountUnrounded string  `json:"amount_unrounded"`
				AmountRounded   string  `json:"amount_rounded"`
				Currency        string  `json:"currency"`
			} `json:"all_inclusive_amount"`
			NetAmount struct {
				Value           float64 `json:"value"`
				AmountUnrounded string  `json:"amount_unrounded"`
				AmountRounded   string  `json:"amount_rounded"`
				Currency        string  `json:"currency"`
			} `json:"net_amount"`
			IncludedTaxesAndChargesAmount struct {
				AmountUnrounded string  `json:"amount_unrounded"`
				Value           float64 `json:"value"`
				AmountRounded   string  `json:"amount_rounded"`
				Currency        string  `json:"currency"`
			} `json:"included_taxes_and_charges_amount"`
			ExcludedAmount struct {
				Currency        string `json:"currency"`
				AmountRounded   string `json:"amount_rounded"`
				AmountUnrounded string `json:"amount_unrounded"`
				Value           int    `json:"value"`
			} `json:"excluded_amount"`
			ChargesDetails struct {
				Amount struct {
					Value    int    `json:"value"`
					Currency string `json:"currency"`
				} `json:"amount"`
				Mode           string `json:"mode"`
				TranslatedCopy string `json:"translated_copy"`
			} `json:"charges_details"`
			ProductPriceBreakdowns []struct {
				GrossAmountHotelCurrency struct {
					Value           float64 `json:"value"`
					AmountUnrounded string  `json:"amount_unrounded"`
					AmountRounded   string  `json:"amount_rounded"`
					Currency        string  `json:"currency"`
				} `json:"gross_amount_hotel_currency"`
				GrossAmount struct {
					AmountRounded   string  `json:"amount_rounded"`
					AmountUnrounded string  `json:"amount_unrounded"`
					Value           float64 `json:"value"`
					Currency        string  `json:"currency"`
				} `json:"gross_amount"`
				ChargesDetails struct {
					Mode           string `json:"mode"`
					TranslatedCopy string `json:"translated_copy"`
					Amount         struct {
						Value    int    `json:"value"`
						Currency string `json:"currency"`
					} `json:"amount"`
				} `json:"charges_details"`
				Benefits []struct {
					Kind         string      `json:"kind"`
					Identifier   string      `json:"identifier"`
					Icon         interface{} `json:"icon"`
					Details      string      `json:"details"`
					Name         string      `json:"name"`
					BadgeVariant string      `json:"badge_variant"`
				} `json:"benefits"`
				GrossAmountPerNight struct {
					Currency        string  `json:"currency"`
					AmountUnrounded string  `json:"amount_unrounded"`
					AmountRounded   string  `json:"amount_rounded"`
					Value           float64 `json:"value"`
				} `json:"gross_amount_per_night"`
				Items []struct {
					Details *string `json:"details"`
					Base    struct {
						BaseAmount int    `json:"base_amount,omitempty"`
						Kind       string `json:"kind"`
						Percentage int    `json:"percentage,omitempty"`
					} `json:"base"`
					InclusionType string `json:"inclusion_type,omitempty"`
					ItemAmount    struct {
						AmountUnrounded string  `json:"amount_unrounded"`
						Value           float64 `json:"value"`
						AmountRounded   string  `json:"amount_rounded"`
						Currency        string  `json:"currency"`
					} `json:"item_amount"`
					Name       string `json:"name"`
					Kind       string `json:"kind"`
					Identifier string `json:"identifier,omitempty"`
				} `json:"items"`
				AllInclusiveAmount struct {
					Currency        string  `json:"currency"`
					AmountUnrounded string  `json:"amount_unrounded"`
					Value           float64 `json:"value"`
					AmountRounded   string  `json:"amount_rounded"`
				} `json:"all_inclusive_amount"`
				ExcludedAmount struct {
					Currency        string `json:"currency"`
					AmountRounded   string `json:"amount_rounded"`
					AmountUnrounded string `json:"amount_unrounded"`
					Value           int    `json:"value"`
				} `json:"excluded_amount"`
				NetAmount struct {
					AmountRounded   string  `json:"amount_rounded"`
					AmountUnrounded string  `json:"amount_unrounded"`
					Value           float64 `json:"value"`
					Currency        string  `json:"currency"`
				} `json:"net_amount"`
				IncludedTaxesAndChargesAmount struct {
					Value           float64 `json:"value"`
					AmountUnrounded string  `json:"amount_unrounded"`
					AmountRounded   string  `json:"amount_rounded"`
					Currency        string  `json:"currency"`
				} `json:"included_taxes_and_charges_amount"`
				StrikethroughAmountPerNight struct {
					Currency        string  `json:"currency"`
					AmountRounded   string  `json:"amount_rounded"`
					AmountUnrounded string  `json:"amount_unrounded"`
					Value           float64 `json:"value"`
				} `json:"strikethrough_amount_per_night,omitempty"`
				StrikethroughAmount struct {
					Value           float64 `json:"value"`
					AmountUnrounded string  `json:"amount_unrounded"`
					AmountRounded   string  `json:"amount_rounded"`
					Currency        string  `json:"currency"`
				} `json:"strikethrough_amount,omitempty"`
				DiscountedAmount struct {
					Currency        string  `json:"currency"`
					AmountUnrounded string  `json:"amount_unrounded"`
					Value           float64 `json:"value"`
					AmountRounded   string  `json:"amount_rounded"`
				} `json:"discounted_amount,omitempty"`
			} `json:"product_price_breakdowns"`
			GrossAmountHotelCurrency struct {
				Currency        string  `json:"currency"`
				AmountUnrounded string  `json:"amount_unrounded"`
				Value           float64 `json:"value"`
				AmountRounded   string  `json:"amount_rounded"`
			} `json:"gross_amount_hotel_currency"`
			DiscountedAmount struct {
				Currency        string  `json:"currency"`
				Value           float64 `json:"value"`
				AmountUnrounded string  `json:"amount_unrounded"`
				AmountRounded   string  `json:"amount_rounded"`
			} `json:"discounted_amount,omitempty"`
			StrikethroughAmount struct {
				AmountRounded   string  `json:"amount_rounded"`
				AmountUnrounded string  `json:"amount_unrounded"`
				Value           float64 `json:"value"`
				Currency        string  `json:"currency"`
			} `json:"strikethrough_amount,omitempty"`
			StrikethroughAmountPerNight struct {
				Value           float64 `json:"value"`
				AmountUnrounded string  `json:"amount_unrounded"`
				AmountRounded   string  `json:"amount_rounded"`
				Currency        string  `json:"currency"`
			} `json:"strikethrough_amount_per_night,omitempty"`
		} `json:"composite_price_breakdown"`
		MainPhotoId            int         `json:"main_photo_id"`
		IsFreeCancellable      int         `json:"is_free_cancellable"`
		Currencycode           string      `json:"currencycode"`
		CantBook               interface{} `json:"cant_book"`
		UnitConfigurationLabel string      `json:"unit_configuration_label"`
		Url                    string      `json:"url"`
		Type                   string      `json:"type"`
		HasFreeParking         int         `json:"has_free_parking,omitempty"`
		HotelHasVbBoost        int         `json:"hotel_has_vb_boost"`
		PriceBreakdown         struct {
			SumExcludedRaw         string      `json:"sum_excluded_raw"`
			AllInclusivePrice      float64     `json:"all_inclusive_price"`
			GrossPrice             interface{} `json:"gross_price"`
			HasIncalculableCharges int         `json:"has_incalculable_charges"`
			HasFinePrintCharges    int         `json:"has_fine_print_charges"`
			HasTaxExceptions       int         `json:"has_tax_exceptions"`
			Currency               string      `json:"currency"`
		} `json:"price_breakdown"`
		HotelId                  int      `json:"hotel_id"`
		IsGeoRate                string   `json:"is_geo_rate"`
		IsNoPrepaymentBlock      int      `json:"is_no_prepayment_block"`
		NativeAdId               string   `json:"native_ad_id"`
		AddressTrans             string   `json:"address_trans"`
		Distance                 string   `json:"distance"`
		ReviewScoreWord          string   `json:"review_score_word"`
		IsGeniusDeal             int      `json:"is_genius_deal"`
		Class                    int      `json:"class"`
		ReviewScore              *float64 `json:"review_score"`
		HotelName                string   `json:"hotel_name"`
		GeniusDiscountPercentage int      `json:"genius_discount_percentage"`
		ReviewRecommendation     string   `json:"review_recommendation"`
		Countrycode              string   `json:"countrycode"`
		Checkin                  struct {
			Until string `json:"until"`
			From  string `json:"from"`
		} `json:"checkin"`
		Latitude              float64     `json:"latitude"`
		Zip                   string      `json:"zip"`
		CityTrans             string      `json:"city_trans"`
		HotelFacilities       string      `json:"hotel_facilities"`
		SelectedReviewTopic   interface{} `json:"selected_review_topic"`
		MaxPhotoUrl           string      `json:"max_photo_url"`
		Max1440PhotoUrl       string      `json:"max_1440_photo_url"`
		HotelIncludeBreakfast int         `json:"hotel_include_breakfast,omitempty"`
		BookingHome           struct {
			IsSingleUnitProperty interface{} `json:"is_single_unit_property"`
			QualityClass         int         `json:"quality_class"`
			Segment              int         `json:"segment"`
			Group                string      `json:"group"`
			IsBookingHome        int         `json:"is_booking_home"`
		} `json:"booking_home,omitempty"`
		ExternalReviews struct {
			NumReviews    int     `json:"num_reviews"`
			ShouldDisplay string  `json:"should_display"`
			Score         float64 `json:"score"`
			ScoreWord     string  `json:"score_word"`
		} `json:"external_reviews,omitempty"`
		HasSwimmingPool int `json:"has_swimming_pool,omitempty"`
	} `json:"result"`
}
