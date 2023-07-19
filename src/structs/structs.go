package structs

type Stats struct {
	ID         string `json:"id"`
	Name       string `json:"fullName"`
	Reputation int    `json:"reputation"`
	Gold       int    `json:"gold"`
	Silver     int    `json:"silver"`
	Bronze     int    `json:"bronze"`
	ImageUrl   string `json:"imageUrl"`
}

type StackStats struct {
	Items []struct {
		DisplayName  string `json:"display_name"`
		ProfileImage string `json:"profile_image"`
		Reputation   int    `json:"reputation"`
		BadgeCounts  struct {
			Bronze int `json:"bronze"`
			Gold   int `json:"gold"`
			Silver int `json:"silver"`
		} `json:"badge_counts"`
	} `json:"items"`
}

type Theme struct {
	TitleColor string `json:"title_color"`
	IconColor  string `json:"icon_color"`
	TextColor  string `json:"text_color"`
	BgColor    string `json:"bg_color"`
}