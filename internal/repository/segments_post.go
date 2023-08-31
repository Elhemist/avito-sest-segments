package repository

import (
	segment "avito-sest-segments/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type SegmentPostgres struct {
	db *sqlx.DB
}

func NewSegmentPostgres(db *sqlx.DB) *SegmentPostgres {
	return &SegmentPostgres{db: db}

}
func (r *SegmentPostgres) CreateSegment(seg segment.Segment) (segment.Segment, error) {
	var userSeg segment.Segment
	if seg == (segment.Segment{}) {
		query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1)", segmentTable)
		_, err := r.db.Exec(query, seg.Name)
		if err != nil {
			return userSeg, fmt.Errorf("Segment isert error")
		}
		query = fmt.Sprintf("SELECT id, name FROM %s WHERE name=$1 ", segmentTable)
		err = r.db.Select(&userSeg, query, seg.Name)
		return userSeg, err
	} else {
		return userSeg, fmt.Errorf("Empty name")
	}
}
