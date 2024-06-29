package tools

import (
	"database/sql"

	"github.com/aixoio/speed-notes/server/data"
)

func AllNotesByUserID(uid uint32, db *sql.DB) ([]data.Note, error) {
	rows, err := db.Query("SELECT * FROM notes WHERE user_id = $1", uid)
	if err != nil {
		return []data.Note{}, err
	}
	defer rows.Close()

	notes := []data.Note{}

	for rows.Next() {
		var id *uint32
		var uid *uint32
		var title *string
		var content *string

		err := rows.Scan(&id, &uid, &title, &content)
		if err != nil {
			return []data.Note{}, err
		}
	}

	return notes, nil
}
