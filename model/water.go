package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	Water struct {
		ID   bson.ObjectId `json:"-" bson:"_id,omitempty"`
		UID  bson.ObjectId `json:"uid,omitempty" bson:"uid,omitempty"`
		Logs []WaterLog    `json:"logs,omitempty" bson:"logs"`
	}

	WaterLog struct {
		Date    time.Time `json:"date,omitempty" bson:"date"`
		WaterIn int       `json:"waterIn" bson:"waterIn"`
	}
)
