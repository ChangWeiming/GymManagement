package member

//Member member info
type Member struct {
	ID          int    `json:"Id"`
	Name        string `json:"Name"`
	Gender      string `json:"Gender"`
	Password    string `json:"Password"`
	Age         int    `json:"Age"`
	PhoneNumber int    `json:"PhoneNumber"`
	Address     string `json:"Address"`
	Term        int    `json:"term"`
}
