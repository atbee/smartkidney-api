package route

import (
	"github.com/labstack/echo"
	"github.com/rocketblack/smartkidney-api/api"
)

// Init initialize api routes and set up a connection.
func Init(e *echo.Echo) {
	// Database connection.
	db, err := api.NewMongoDB()
	if err != nil {
		e.Logger.Fatal(err)
	}

	a := &api.MongoDB{
		Conn:     db.Conn,
		UCol:     db.UCol,
		BPCol:    db.BPCol,
		GIRCol:   db.GIRCol,
		BSCol:    db.BSCol,
		BMICol:   db.BMICol,
		WATERCol: db.WATERCol,
	}

	// Routes
	// Authentication.
	e.POST("/login", a.Login)
	e.POST("/register", a.Register)

	// Users.
	e.GET("/users/:id", a.FindUser)
	e.PATCH("/users/:id", a.EditUser)
	e.DELETE("/users/:id", a.DeleteUser)

	// Blood pressures(BP).
	e.GET("/bp/:id", a.ViewBP)
	e.POST("/bp/:id", a.AddBP)

	// Glomerular infiltration rates(GIR).
	e.GET("/gir/:id", a.ViewGIR)
	e.POST("/gir/:id", a.AddGIR)

	// Blood sugar(BS).
	e.GET("/bs/:id", a.ViewBS)
	e.POST("/bs/:id", a.AddBS)

	// BMI
	e.GET("/bmi/:id", a.ViewBMI)
	e.POST("/bmi/:id", a.AddBMI)

	// Water
	e.GET("/bmi/:id", a.ViewWater)
	e.POST("/bmi/:id", a.AddWater)
}
