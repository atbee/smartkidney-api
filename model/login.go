package model

// LoginRes holds metadata about login response.
type (
	LoginRes struct {
		FirstLogin bool `json:"firstLogin"`
		User       User `json:"user,omitempty" bson:"user,omitempty"`
	}
)
