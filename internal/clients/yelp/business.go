package yelp

// Business describes a business on yelp
// https://www.yelp.com/developers/documentation/v3/business
type Business struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ImageURL    string `json:"image_url"`
	IsClosed    bool   `json:"is_closed"`
	Phone       string `json:"phone"`
	ReviewCount int64  `json:"review_count"`
	Categories  []struct {
		Alias string `json:"alias"`
		Title string `json:"title"`
	} `json:"categories"`
	Rating   float64 `json:"rating"`
	Location struct {
		Address1       string   `json:"address1"`
		Address2       string   `json:"address2"`
		Address3       string   `json:"address3"`
		City           string   `json:"city"`
		ZipCode        string   `json:"zip_code"`
		Country        string   `json:"country"`
		State          string   `json:"state"`
		DisplayAddress []string `json:"display_address"`
		CrossStreets   string   `json:"cross_streets"`
	} `json:"location"`
	Coordinates struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinates"`
	Photos []string `json:"photos"`
	Price  string   `json:"price"`
	Hours  []struct {
		Open []struct {
			IsOvernight bool   `json:"is_overnight"`
			Start       string `json:"start"`
			End         string `json:"end"`
			Day         int64  `json:"day"`
		} `json:"open"`
		HoursType string `json:"hours_type"`
		IsOpenNow bool   `json:"is_open_now"`
	} `json:"hours"`
	Transactions []string `json:"transactions"`
	SpecialHours []struct {
		Date        string `json:"date"`
		IsClosed    bool   `json:"is_closed"`
		Start       string `json:"start"`
		End         string `json:"end"`
		IsOvernight bool   `json:"is_overnight"`
	} `json:"special_hours"`
}
