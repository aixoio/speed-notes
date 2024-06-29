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

		notes = append(notes, data.Note{
			Id:       *id,
			User_id:  *uid,
			Title:    *content,
			Contents: *content,
		})
	}

	return notes, nil
}

func InsertNewNote(uid uint32, title, content string, db *sql.DB) (uint32, error) {
	var note_id *uint32
	err := db.QueryRow("INSERT INTO notes(user_id, title, contents) VALUES ($1, $2, $3) RETURNING id", uid, title, content).Scan(&note_id)
	if err != nil {
		return 0, err
	}
	return *note_id, nil
}

func NoteByID(id uint32, db *sql.DB) (data.Note, error) {
	var note_id *uint32
	var user_id *uint32
	var title *string
	var contents *string
	err := db.QueryRow("SELECT * FROM notes WHERE id = $1", id).Scan(&note_id, &user_id, &title, &contents)
	if err != nil {
		return data.Note{}, err
	}
	return data.Note{
		Id:       *note_id,
		User_id:  *user_id,
		Title:    *title,
		Contents: *contents,
	}, nil
}
