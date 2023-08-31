package repository

import (
	segment "avito-sest-segments/models"

	"github.com/jmoiron/sqlx"
)

type ActiveSegmentPostgres struct {
	db *sqlx.DB
}

func NewASegPostgres(db *sqlx.DB) *ActiveSegmentPostgres {
	return &ActiveSegmentPostgres{db: db}
}

//	type ActiveSegment interface {
//		AddActiveSegment(segment.User) (segment.User, error)
//		DeleteActiveSegment(segment.User) (segment.User, error)
//	}
func (r *ActiveSegmentPostgres) AddActiveSegment(user segment.ActiveSegment) (segment.ActiveSegment, error) {

	return user, nil
}
func (r *ActiveSegmentPostgres) DeleteActiveSegment(user segment.ActiveSegment) error {
	return nil
}
