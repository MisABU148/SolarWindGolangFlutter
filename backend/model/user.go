package model

type UserDTO struct {
	ID               int64   `json:"id"`
	TelegramID       int64   `json:"telegramId"`
	Username         string  `json:"username"`
	Description      string  `json:"description"`
	Age              int64   `json:"age"`
	Gender           string  `json:"gender"`
	Verified         bool    `json:"verified"`
	PreferredGender  string  `json:"preferredGender"`
	CityID           int64   `json:"cityId"`
	PreferredGymTime []int   `json:"preferredGymTime"`
	SportIDs         []int64 `json:"sportId"`
}

type User struct {
	ID               int64
	TelegramID       int64
	UserName         string
	Description      string
	Age              int64
	Gender           string
	Verified         bool
	PreferredGender  string
	CityID           int64
	PreferredGymTime int
	SportIDs         []int64
}
