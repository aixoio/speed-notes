package tools

import (
	"database/sql"
	"regexp"

	"github.com/aixoio/speed-notes/server/data"
	"golang.org/x/crypto/bcrypt"
)

const BCRYPT_COST = 14

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

func UserFromUsername(username string, db *sql.DB) (data.User, error) {

	row := db.QueryRow("SELECT * FROM users WHERE username = $1", username)

	var id *uint32
	var username_db *string
	var password *string

	err := row.Scan(&id, &username_db, &password)
	if err != nil {
		return data.User{}, err
	}

	return data.User{Id: *id, Username: *username_db, Password_hash: *password}, nil
}

func UserInsert(user data.User, db *sql.DB) error {

	_, err := db.Exec("INSERT INTO users(username, password_hash) VALUES ($1, $2)", user.Username, user.Password_hash)

	return err
}

func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), BCRYPT_COST)
	return string(hash), err
}

func VerifyPassword(hash string, pwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd)) == nil
}
