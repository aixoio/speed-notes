package tools

import (
	"database/sql"
	"regexp"
)

func IsUsernameValid(un string) bool {
	return !regexp.MustCompile("[^A-z0-9]+").MatchString(un)
}

func UserExists(username string, db *sql.DB) (bool, error) {
	if !IsUsernameValid(username) {
		return false, nil
	}

}
