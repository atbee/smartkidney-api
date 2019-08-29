package api

import (
	"github.com/rocketblack/smartkidney-api/model"
)

// FindUser search and validation of user data.
// func (db *MongoDB) FindUser() (*model.User, string) {
// 	u := &model.User{}

// 	q := bson.M{
// 		"email": u.Email,
// 	}
// 	if err := db.UCol.Find(q).One(&u); err != nil {
// 		return nil, false
// 	}

// 	return u, status
// }

// CreateUser created by ID and email in database.
func (db *MongoDB) CreateUser(u *model.User) {
	if err := db.UCol.Insert(&u); err != nil {
		return
	}
}
