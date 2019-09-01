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
		Conn:  db.Conn,
		UCol:  db.UCol,
		BPCol: db.BPCol,
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

	// Blood sugar(BS).
}
