package repository

import (
	segment "avito-sest-segments/models"

	"github.com/jmoiron/sqlx"
)

type SegmentPostgres struct {
	db *sqlx.DB
}

func NewSegmentPostgres(db *sqlx.DB) *SegmentPostgres {
	return &SegmentPostgres{db: db}

}

func (r *SegmentPostgres) CreateSegment(thisSegment segment.Segment) (int, error) {

	return 0, nil
}
