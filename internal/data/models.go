package data

import (
	"database/sql"
	"errors"
)

// Define a custom ErrRecordNotFound error.
// We'll return this from our Get() method when looking up a movie that doesn't exist in our database.
var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Tasks       TaskModel
	Permissions PermissionModel // Add a new Permissions field.
	Tokens      TokenModel      // Add a new Tokens field.
	Users       UserModel       // Add a new Users field.
}

// For ease of use, we also add a New() method which returns a Models struct containing the initialized MovieModel.
func NewModels(db *sql.DB) Models {
	return Models{
		Tasks:       TaskModel{DB: db},
		Permissions: PermissionModel{DB: db}, // Initialize a new PermissionModel instance.
		Tokens:      TokenModel{DB: db},      // Initialize a new TokenModel instance.
		Users:       UserModel{DB: db},       // Initialize a new UserModel instance.
	}
}
