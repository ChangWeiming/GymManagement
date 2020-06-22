package member

//Member member info
type Member struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Password    string `json:"password"`
	Age         string `json:"age"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Term        string `json:"term"`
}
