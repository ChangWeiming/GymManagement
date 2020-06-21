package member

//Member member info
type Member struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Password    string `json:"password"`
	Age         int    `json:"age"`
	PhoneNumber int    `json:"phone_number"`
	Address     string `json:"address"`
	Term        int    `json:"term"`
}
