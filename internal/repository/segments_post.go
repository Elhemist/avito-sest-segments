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
func (r *SegmentPostgres) CreateSegment(name string) (int, error) {

	var userSeg segment.Segment
	if name != "" {
		query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1)", segmentTable)
		_, err := r.db.Exec(query, name)
		if err != nil {
			return userSeg.Id, fmt.Errorf("Segment insert error")
		}
		query = fmt.Sprintf("SELECT id, name FROM %s WHERE name=$1 ", segmentTable)
		err = r.db.Get(&userSeg, query, name)
		return userSeg.Id, err
	} else {
		return userSeg.Id, fmt.Errorf("Empty name")
	}
}
func (r *SegmentPostgres) DeleteSegment(name string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE name = $1", segmentTable)
	_, err := r.db.Exec(query, name)

	return err
}
