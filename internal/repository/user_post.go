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

//	type User interface {
//		AddUser(segment.User) (segment.User, error)
//		CheckUser(segment.User) (segment.User, error)
//	}
func (r *UserPostgres) AddUser(user segment.User) (segment.User, error) {
	var userRes segment.User

	return userRes, nil
}
func (r *UserPostgres) CheckUser(user segment.User) (segment.User, error) {
	var userRes segment.User

	query := fmt.Sprintf("SELECT id, balance FROM %s WHERE id=$1 ", "users")
	err := r.db.Get(&userRes, query, user.Id)
	return userRes, err
}
