package repository

import (
	segment "avito-sest-segments/models"

	"github.com/jmoiron/sqlx"
)

type Segment interface {
	CreateSegment(string) (int, error)
	DeleteSegment(string) error
}
type User interface {
	AddUser(segment.User) error
	ExistUser(userId int) error
	CheckUser(segment.User) ([]int, error)
	AddSegments(segId int, userId int, expire string) error
	DeleteActiveSegment(serId int, userId int) error
	GetSegmentByName(string) (int, error)
	GetSegmentById(id int) (string, error)
	GetPartUsers(part int) ([]int, error)
}
type ActiveSegment interface {
	AddActiveSegment(segment.ActiveSegment) (segment.ActiveSegment, error)
	DeleteActiveSegment(segment.ActiveSegment) error
}
type Repository struct {
	Segment
	User
	ActiveSegment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Segment:       NewSegmentPostgres(db),
		User:          NewUserPostgres(db),
		ActiveSegment: NewASegPostgres(db),
	}
}
