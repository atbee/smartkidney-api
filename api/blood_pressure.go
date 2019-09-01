package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/rocketblack/smartkidney-api/model"
	"gopkg.in/mgo.v2/bson"
)

// ViewBP show logs of blood pressure for each user.
// You can view a list of blood pressure each week by query parameter.
// &week=00&year==0000
func (db *MongoDB) ViewBP(c echo.Context) error {
	id := c.Param("id")
	bp := []*model.BloodPressureLog{}

	match := bson.M{
		"$match": bson.M{
			"uid": bson.ObjectIdHex(id),
		},
	}

	unwind := bson.M{
		"$unwind": "$logs",
	}

	replaceRoot := bson.M{
		"$replaceRoot": bson.M{
			"newRoot": "$logs",
		},
	}

	sort := bson.M{
		"$sort": bson.M{
			"date": 1,
		},
	}

	q := []bson.M{match, unwind, replaceRoot, sort}

	w := c.QueryParam("week")
	y := c.QueryParam("year")
	week, _ := strconv.Atoi(w)
	year, _ := strconv.Atoi(y)

	if w != "" && y != "" {
		set := bson.M{
			"$set": bson.M{
				"week": bson.M{"$week": "$date"},
				"year": bson.M{"$year": "$date"},
			},
		}

		matchWeek := bson.M{
			"$match": bson.M{
				"week": week,
				"year": year,
			},
		}

		q = []bson.M{match, unwind, replaceRoot, set, matchWeek, sort}
	}

	if err := db.BPCol.Pipe(q).All(&bp); err != nil {
		return c.JSON(http.StatusNotFound, "the user not found.")
	}

	if len(bp) == 0 {
		return c.JSON(http.StatusNotFound, "user data is empty or not found.")
	}

	return c.JSON(http.StatusOK, &bp)
}

// AddBP add items systolic and diastolic to that user.
func (db *MongoDB) AddBP(c echo.Context) error {
	id := c.Param("id")
	bp := &model.BloodPressureLog{
		Date: time.Now(),
	}
	if err := c.Bind(bp); err != nil {
		return err
	}

	q := bson.M{
		"uid": bson.ObjectIdHex(id),
	}

	d := bson.M{
		"$push": bson.M{
			"logs": bp,
		},
	}

	if err := db.BPCol.Update(q, &d); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, &bp)
}

// CreateBP created logs of blood pressure in database.
func (db *MongoDB) CreateBP(uid bson.ObjectId) {
	bp := &model.BloodPressure{
		ID:  bson.NewObjectId(),
		UID: uid,
	}
	if err := db.BPCol.Insert(&bp); err != nil {
		return
	}
}
