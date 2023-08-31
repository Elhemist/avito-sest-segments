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
	if user == (segment.User{}) {
		query := fmt.Sprintf("INSERT INTO %s (id) VALUES ($1)", userTable)
		_, err := r.db.Exec(query, user.Id)
		return err
	} else {
		return fmt.Errorf("Empty name")
	}
}
func (r *UserPostgres) CheckUser(user segment.User) ([]segment.Segment, error) {
	var userSeg []segment.Segment

	query := fmt.Sprintf("SELECT id, balance FROM %s WHERE id=$1 ", userTable)
	err := r.db.Select(&userSeg, query, user.Id)
	return userSeg, err
}
