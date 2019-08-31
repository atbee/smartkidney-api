package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User holds metadata about a user.
type (
	User struct {
		ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
		Name      string        `json:"name,omitempty" bson:"name" validate:"nonzero"`
		Email     string        `json:"email,omitempty" bson:"email" validate:"nonzero"`
		BirthDate string        `json:"birthDate,omitempty" bson:"birthDate" validate:"len=10"`
		Gender    string        `json:"gender,omitempty" bson:"gender" validate:"regexp=(^male$|^female$)"`
		Hospital  string        `json:"hospital,omitempty" bson:"hospital" validate:"nonzero"`
		Weight    string        `json:"weight,omitempty" bson:"weight"`
		Height    string        `json:"height,omitempty" bson:"height"`
		CreateAt  time.Time     `json:"-" bson:"createAt"`
	}
)
