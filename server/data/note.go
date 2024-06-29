package data

type Note struct {
	Id       uint32 `json:"id"`
	User_id  uint32 `json:"user_id"`
	Title    string `json:"title"`
	Contents string `json:"contents"`
}
