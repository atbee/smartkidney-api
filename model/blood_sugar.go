package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	BloodSugar struct {
		ID   bson.ObjectId   `json:"-" bson:"_id,omitempty"`
		UID  bson.ObjectId   `json:"uid,omitempty" bson:"uid,omitempty"`
		Logs []BloodSugarLog `json:"logs,omitempty" bson:"logs"`
	}

	BloodSugarLog struct {
		Date       time.Time `json:"date,omitempty" bson:"date"`
		SugarLevel int       `json:"sugarLevel" bson:"sugarLevel"`
		HbA1c      int       `json:"hba1c" bson:"hba1c"`
	}
)
