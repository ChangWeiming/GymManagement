package course

//Course course info
type Course struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Time         string `json:"time"`
	PeopleNumber string `json:"people_number"`
}
