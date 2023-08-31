package segment

type User struct {
	Id int `json:"userId"`
}

type SegmentsToUpdate struct {
	Add    []string `json:"segToAdd"`
	Delete []string `json:"segToDelete"`
	UserId int      `json:"userId"`
}
