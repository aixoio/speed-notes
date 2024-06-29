package tools

import (
	"database/sql"
	"regexp"

	"github.com/aixoio/speed-notes/server/data"
)

func IsUsernameValid(un string) bool {
	return !regexp.MustCompile("[^A-z0-9]+").MatchString(un)
}

func UserExists(username string, db *sql.DB) (bool, error) {

	row := db.QueryRow("SELECT id FROM users WHERE username = $1", username)

	var _id *uint = nil

	err := row.Scan(&_id)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func UserInsert(user data.User, db *sql.DB) error {

	return nil
}
