package todo

const (
	FIND_TODOS_BY_USER_ID_SQL = "SELECT * FROM todos WHERE user_id=$1"
)
