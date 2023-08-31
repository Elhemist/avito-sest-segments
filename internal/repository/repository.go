package repository

import (
	segment "avito-sest-segments/models"

	"github.com/jmoiron/sqlx"
)

type Segment interface {
	CreateSegment(segment.Segment) (int, error)
}
type User interface {
	AddUser(segment.User) (segment.User, error)
	CheckUser(segment.User) (segment.User, error)
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
