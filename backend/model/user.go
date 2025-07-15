package model

type UserDTO struct {
	ID               int64  `json:"id"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Description      string `json:"description"`
	Age              int    `json:"age"`
	PreferredGender  string `json:"preferredgender"`
	Gender           string `json:"gender"`
	CityId           int    `json:"cityId"`
	PreferredGymTime []int  `json:"preferredGymTime"`
	SportId          []int  `json:"sportId"`
}

type User struct {
	ID               int64
	UserName         string
	PasswordHash     string
	Alias            string
	Description      string
	Age              int
	PreferredGender  string
	Gender           string
	CityId           int
	PreferredGymTime []string
	SportId          []string
}
