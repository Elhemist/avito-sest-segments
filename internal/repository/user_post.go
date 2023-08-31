package repository

import (
	segment "avito-sest-segments/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) AddUser(user segment.User) error {
	if user != (segment.User{}) {
		query := fmt.Sprintf("INSERT INTO %s (id) VALUES ($1)", userTable)
		_, err := r.db.Exec(query, user.Id)
		return err
	} else {
		return fmt.Errorf("Empty name")
	}
}
func (r *UserPostgres) CheckUser(user segment.User) ([]segment.Segment, error) {

	//todo!!
	var activeSeg []int
	var userSeg []segment.Segment

	query := fmt.Sprintf("SELECT segmentId FROM activeSegments WHERE userId=$1 ")
	err := r.db.Select(&activeSeg, query, user.Id)
	if err != nil {
		for _, x := range activeSeg {
			query = fmt.Sprintf("SELECT id, name FROM %s WHERE id=$1 ", segmentTable)
			err = r.db.Select(&activeSeg, query, x)
		}
	}
	//todo!!
	return userSeg, err
}
func (r *UserPostgres) AddSegments(user segment.User, servise segment.Segment) (int, error) {
	if user != (segment.User{}) {
		query := fmt.Sprintf("INSERT INTO %s (id) VALUES ($1)", userTable)
		_, err := r.db.Exec(query, user.Id)
		return 0, err
	} else {
		return 0, fmt.Errorf("Empty name")
	}
}
