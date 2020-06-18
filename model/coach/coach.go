package coach

//Coach coach info
type Coach struct {
	Name            string `json:"name"`
	ID              int    `json:"id"`
	Gender          string `json:"gender"`
	Password        string `json:"password"`
	Age             int    `json:"age"`
	PhoneNumber     int    `json:"phone_number"`
	Address         string `json:"address"`
	PersonalProfile string `json:"personal_profile"`
}
