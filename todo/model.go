package todo

type Todo struct {
	ID          string `json:"id"`
	UserID      string `db:"user_id" json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Content     string `json:"content,omitempty"`
}
