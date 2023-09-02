package segment

type User struct {
	Id int `json:"userId"`
}
type SegmentAdd struct {
	Name       string `json:"segName"`
	ExpireTime string `json:"segDeleteTime"`
}
type SegmentsToUpdate struct {
	Add    []SegmentAdd `json:"segToAdd"`
	Delete []string     `json:"segToDelete"`
	UserId int          `json:"userId"`
}

type SegmentsPartAdd struct {
	Part       int    `json:"part"`
	SegName    string `json:"segName"`
	ExpireTime string `json:"segDeleteTime"`
}
