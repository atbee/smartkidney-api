package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	GlomerularInfil struct {
		ID   bson.ObjectId        `json:"-" bson:"_id,omitempty"`
		UID  bson.ObjectId        `json:"uid,omitempty" bson:"uid,omitempty"`
		Logs []GlomerularInfilLog `json:"logs,omitempty" bson:"logs"`
	}

	GlomerularInfilLog struct {
		Date time.Time `json:"date,omitempty" bson:"date"`
		Cr   float64   `json:"cr" bson:"cr"`
		EGFR float64   `json:"egfr" bson:"egfr"`
	}
)
