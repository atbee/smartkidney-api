package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	BMI struct {
		ID   bson.ObjectId `json:"-" bson:"_id,omitempty"`
		UID  bson.ObjectId `json:"uid,omitempty" bson:"uid,omitempty"`
		Logs []BMILog      `json:"logs,omitempty" bson:"logs"`
	}

	BMILog struct {
		Date time.Time `json:"date,omitempty" bson:"date"`
		BMI  float64   `json:"bmi" bson:"bmi"`
	}
)
