package structs

type Stats struct {
	ID         string `json:"id"`
	Name       string `json:"fullName"`
	Reputation int    `json:"reputation"`
	Gold       int    `json:"gold"`
	Silver     int    `json:"silver"`
	Bronze     int    `json:"bronze"`
	Questions  int    `json:"question_count"`
	Answers    int    `json:"answer_count"`
	LastSeen   int64  `json:"last_access_date"`
	ViewCount  int    `json:"view_count"`
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
		Questions int   `json:"question_count"`
		Answers   int   `json:"answer_count"`
		LastSeen  int64 `json:"last_access_date"`
		ViewCount int   `json:"view_count"`
	} `json:"items"`
}

type Theme struct {
	Gold      string `json:"gold"`
	Silver    string `json:"silver"`
	Bronze    string `json:"bronze"`
	BgColor   string `json:"bg_color"`
	TextColor string `json:"text_color"`
}