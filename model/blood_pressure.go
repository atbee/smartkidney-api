package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// BloodPressure is
	BloodPressure struct {
		ID   bson.ObjectId      `json:"-" bson:"_id,omitempty"`
		UID  bson.ObjectId      `json:"uid,omitempty" bson:"uid,omitempty"`
		Logs []BloodPressureLog `json:"logs,omitempty" bson:"logs"`
	}
	// BloodPressureLog is
	BloodPressureLog struct {
		Date      time.Time `json:"date,omitempty" bson:"date"`
		Systolic  int       `json:"systolic" bson:"systolic"`
		Diastolic int       `json:"diastolic" bson:"diastolic"`
	}
)
